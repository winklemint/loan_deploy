package control

import (
	con "admin/Config"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB
var jwtKey = []byte("secret key")

type AdminLogin struct {
	ID       int    `json:"id"`
	Name     string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Contact  int    `json:"contact"`
	Role     string `json:"role"`
}

type Credential struct {
	Name     string `json:"name,omitempty"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Contact  int    `json:"contact"`
	jwt.StandardClaims
}

type Session struct {
	ID         int       `json:"id,omitempty"`
	AdminID    int       `json:"admin_id,omitempty"`
	Token      string    `json:"token,omitempty"`
	ExpiresAt  time.Time `json:"expires_at,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
	Ip_address string    `json:"ip_address,omitempty"`
}

//var sessionCookieName = "session_id"

//var store = sessions.NewCookieStore([]byte("your-secret-key"))

// type TemplateData struct {
// 	Session  *Session
// 	Data     string
// 	Username string
// }

//var templates *template.Template

// var templates *template.Template

func init() {
	var err error
	db, err = con.GetDB()
	if err != nil {
		panic(fmt.Errorf("failed to initialize database: %v", err))
	}
}

func AdminInsert(w http.ResponseWriter, r *http.Request) {
	// db, err := con.GetDB()
	// if err != nil {
	// 	http.Error(w, "failed to connect db", http.StatusInternalServerError)
	// 	return
	// }

	// Decode the JSON data from the request body into the AdminLogin struct
	var admin AdminLogin
	err := json.NewDecoder(r.Body).Decode(&admin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	maxPasswordLength := 120 // Modify this value based on your database column's maximum length

	if len(admin.Password) > maxPasswordLength {
		admin.Password = admin.Password[:maxPasswordLength]
	}

	// Encrypt the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	hashedPasswordStr := hex.EncodeToString(hashedPassword)

	// Insert data into the database using a prepared SQL statement
	_, err = db.Exec("INSERT INTO admin(username, password, contact, email, role) VALUES(?, ?, ?, ?, ?)", admin.Name, hashedPasswordStr, admin.Contact, admin.Email, admin.Role)
	if err != nil {
		http.Error(w, "failed to insert data", http.StatusInternalServerError)
		//fmt.Println(err)
		return
	}

	// Send a success JSON response
	type Response struct {
		Data string
	}
	response := Response{
		Data: "Insert Successfully",
	}
	jsonResponse, _ := json.Marshal(response)
	w.Write(jsonResponse)
}

func UpadteAdmin(w http.ResponseWriter, r *http.Request) {
	// db, err := con.GetDB()
	// if err != nil {
	// 	panic(fmt.Errorf("failed to initialize database: %v", err))
	// }
	params := mux.Vars(r)
	AdminID := params["id"]
	var admin AdminLogin
	//fmt.Println("r.Body ================== >", r.Body)
	err := json.NewDecoder(r.Body).Decode(&admin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	maxPasswordLength := 120 // Modify this value based on your database column's maximum length

	if len(admin.Password) > maxPasswordLength {
		admin.Password = admin.Password[:maxPasswordLength]
	}

	// Encrypt the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	hashedPasswordStr := hex.EncodeToString(hashedPassword)
	result, err := db.Exec("UPDATE admin SET username=?,password=?,email=?,contact=? WHERE id=?", admin.Name, hashedPasswordStr, admin.Email, admin.Contact, AdminID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//fmt.Println(rowsAffected)

	if rowsAffected == 0 {
		http.Error(w, "admin not found", http.StatusNotFound)
		return
	}

}
func AdminLoginHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")

	// Handle preflight OPTIONS requests
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Check request method
	if r.Method != http.MethodPost {
		// http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var credential Credential
	if err := json.NewDecoder(r.Body).Decode(&credential); err != nil {
		// http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	admin, err := GetUserFromDatabase(credential.Name)
	if err != nil || admin == nil || admin.Name != credential.Name || !ComparePasswords(credential.Password, admin.Password) {
		logrus.Warnln("Invalid credentials for user:", credential.Name)
		logrus.Warnln("Invalid credentials for admin:", credential.Name)
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	existingSession, err := GetSessionByUserID(admin.ID)
	if err != nil {
		logrus.Infoln(err, "jkdkdfkjksd")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//fmt.Println(existingSession)
	if existingSession != nil {
		// Update the existing session and send the response
		err = UpdateSessionAndRespond(w, existingSession)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	err = CreateSessionAndRespond(w, admin, r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

}
func GetSessionByUserID(adminID int) (*Session, error) {
	// Assuming you have a database connection established
	// db, _ := con.GetDB()

	query := "SELECT token,admin_id FROM admin_sessions WHERE admin_id = ?"
	var session Session

	err := db.QueryRow(query, adminID).Scan(&session.Token, &session.AdminID)
	if err != nil {
		if err == sql.ErrNoRows {
			logrus.Infoln(err, "jkdkdfkjksd")
			return nil, nil // No session found for the user
		}
		return nil, err
	}

	return &session, nil
}
func UpdateSessionAndRespond(w http.ResponseWriter, session *Session) error {
	// Update the session details as needed
	expirationTime := time.Now().Add(1 * time.Minute)
	session.ExpiresAt = expirationTime
	session.UpdatedAt = time.Now()
	session.Token = GenerateNewToken()

	// Update the session in the database
	err := UpdateSession(session)
	if err != nil {
		return err
	}

	// Set the token in the response cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   session.Token,
		Expires: expirationTime,
	})
	adminID := strconv.Itoa(session.AdminID)
	//fmt.Println(adminID)
	http.SetCookie(w, &http.Cookie{
		Name:    "id",
		Value:   adminID,
		Expires: expirationTime,
	})
	w.WriteHeader(http.StatusOK)
	logrus.Infoln("Login Successful", session.AdminID)

	//go removeExpiredSessions()

	return nil
}

func GenerateNewToken() string {
	// Generate a new token as needed
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(1 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(jwtKey)
	return tokenString
}
func UpdateSession(session *Session) error {
	// Assuming you have a database connection established
	// db, _ := con.GetDB()

	query := "UPDATE admin_sessions SET token = ?, expires_at = ?, updated_at = ? WHERE admin_id = ?"
	_, err := db.Exec(query, session.Token, session.ExpiresAt, session.UpdatedAt, session.AdminID)
	if err != nil {
		return err
	}

	return nil
}
func CreateSessionAndRespond(w http.ResponseWriter, admin *AdminLogin, r *http.Request) error {
	expirationTime := time.Now().Add(1 * time.Minute)

	//fmt.Println(expirationTime)

	var c Credential
	claims := &Credential{
		Email: c.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	session := &Session{
		AdminID:   admin.ID,
		Token:     tokenString,
		ExpiresAt: expirationTime,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	adminID := strconv.Itoa(session.AdminID)

	http.SetCookie(w, &http.Cookie{
		Name:    "id",
		Value:   adminID,
		Expires: expirationTime,
	})

	ipAddress := getClientIP(r)

	//fmt.Println(ipAddress)

	err = SaveSession(session, ipAddress)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	logrus.Infoln("Login Successful", admin.Email)

	return nil
}

func removeExpiredSessions() {
	for {
		// Check for expired sessions and delete them
		err := DeleteExpiredSessions()
		if err != nil {
			logrus.Errorln("Error deleting expired sessions:", err)
		}

		//fmt.Println("grtn")

		// Sleep for a certain interval before checking again
		// Adjust the interval based on your application needs
		time.Sleep(12 * time.Hour)
	}
}
func DeleteExpiredSessions() error {
	// Connect to the database
	// db, err := con.GetDB()
	// if err != nil {
	// 	logrus.Errorln("failed to connect to the database:")
	// 	return fmt.Errorf("failed to connect to the database: %v", err)
	// }
	// defer db.Close()

	// Retrieve the expired session tokens from the database
	rows, err := db.Query("SELECT token FROM sessions WHERE  DATE_ADD(created_at, INTERVAL 30 MINUTE) < NOW()")
	if err != nil {
		logrus.Errorln("failed to fetch expired sessions from the database:")
		return fmt.Errorf("failed to fetch expired sessions from the database: %v", err)
	}
	defer rows.Close()

	// Iterate over the rows and delete the corresponding sessions
	for rows.Next() {
		var sessionToken string
		err := rows.Scan(&sessionToken)
		if err != nil {
			logrus.Errorln("failed to scan session token from the database::")
			return fmt.Errorf("failed to scan session token from the database: %v", err)
		}

		// Delete the expired session from the database
		_, err = db.Exec("DELETE FROM sessions WHERE token = ?", sessionToken)
		if err != nil {
			logrus.Errorln("failed to delete expired session from the database")
			return fmt.Errorf("failed to delete expired session from the database: %v", err)
		} else {
			logrus.Errorln("Delete successful")
		}
	}

	return nil
}
func ComparePasswords(passwordHash, hashedPassword string) bool {
	hashBytes, err := hex.DecodeString(hashedPassword)
	if err != nil {
		return false
	}

	err = bcrypt.CompareHashAndPassword(hashBytes, []byte(passwordHash))

	return err == nil
}

// Additional functions omitted for brevity

func GetUserFromDatabase(name string) (*AdminLogin, error) {

	// db, err := con.GetDB()
	// if err != nil {
	// 	panic(fmt.Errorf("failed to initialize database: %v", err))
	// }

	admin := &AdminLogin{}
	query := "SELECT id, username, password, email, contact FROM admin WHERE username = ?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare the query: %v", err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(name).Scan(&admin.ID, &admin.Name, &admin.Password, &admin.Email, &admin.Contact)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to fetch user from the database: %v", err)
	}

	return admin, nil
}

func SaveSession(session *Session, ipAddress string) error {
	// db, err := con.GetDB()
	// if err != nil {
	// 	panic(fmt.Errorf("failed to initialize database: %v", err))
	// }

	query := "INSERT INTO admin_sessions (admin_id, token, expires_at, created_at, updated_at, ip_address) VALUES (?, ?, ?, NOW(), NOW(), ?)"

	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare the query: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(session.AdminID, session.Token, session.ExpiresAt, ipAddress)
	if err != nil {
		//fmt.Println(err)
		return fmt.Errorf("failed to save session in the database: %v", err)
	}

	return nil
}

func GetAdmin(w http.ResponseWriter, r *http.Request) {
	// db, err := con.GetDB()
	// if err != nil {
	// 	panic(fmt.Errorf("failed to initialize database: %v", err))
	// }

	rows, err := db.Query("SELECT id,username,email,password,contact,role FROM admin Where role='Sub Admin'")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var admins []AdminLogin
	var admin AdminLogin

	for rows.Next() {
		err = rows.Scan(&admin.ID, &admin.Name, &admin.Email, &admin.Password, &admin.Contact, &admin.Role)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		admins = append(admins, admin)
	}

	response := Response{
		Data: admins,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
func AdminById(w http.ResponseWriter, r *http.Request) {
	// db, err := con.GetDB()
	// if err != nil {
	// 	panic(fmt.Errorf("failed to initialize database: %v", err))
	// }
	params := mux.Vars(r)
	AdminID := params["id"]
	rows, err := db.Query("SELECT id,username,email,password,contact,role FROM admin Where id=?", AdminID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var admins []AdminLogin
	var admin AdminLogin

	for rows.Next() {
		err = rows.Scan(&admin.ID, &admin.Name, &admin.Email, &admin.Password, &admin.Contact, &admin.Role)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		admins = append(admins, admin)
	}

	response := Response{
		Data: admins,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

type Error struct {
	Message string `json:"message"`
}

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error *Error      `json:"error,omitempty"`
}

func GetAdminByID(w http.ResponseWriter, r *http.Request) {

	// db, err := con.GetDB()
	// if err != nil {
	// 	logrus.Error("Failed to connect to the database:", err)
	// 	// http.Error(w, "Failed to connect to the database", http.StatusInternalServerError)
	// 	return
	// }

	// Set the appropriate headers to enable CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent, Authorization")

	// Retrieve the JWT token from the request header
	authHeader := r.Header.Get("Cookie")

	//fmt.Println(authHeader)

	if authHeader == "" {
		logrus.Warnln("Authorization header missing")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&Response{Error: &Error{Message: "Unauthorized"}})
		return
	}

	// Extract the token from the Authorization header
	tokenString := strings.TrimPrefix(authHeader, "token=")
	tokenString = strings.Split(tokenString, ";")[0]
	//fmt.Println(tokenString)

	session := &Session{
		Token: tokenString,
	}

	// Use the GetUserIDFromSession function to retrieve the userID
	adminID, err := GetUserIDFromSession1(session)
	//fmt.Println(adminID)

	if err != nil {
		logrus.Errorf("Error retrieving adminID from session: %v", err)
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&Response{Error: &Error{Message: "Unauthorized"}})
		return
	}

	var admin AdminLogin

	rows, err := db.Query("SELECT id,username,password,email,contact,role FROM admin WHERE id=?", adminID)
	if err != nil {
		logrus.Errorf("Error executing database query: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var admins []AdminLogin

	for rows.Next() {
		err = rows.Scan(&admin.ID, &admin.Name, &admin.Password, &admin.Email, &admin.Contact, &admin.Role)
		if err != nil {

			logrus.Errorf("Error scanning database rows: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		admins = append(admins, admin)
	}

	jsonResponse, err := json.Marshal(admins)
	if err != nil {
		logrus.Errorf("Error marshalling JSON response: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func GetUserIDFromSession1(session *Session) (int, error) {

	// Validate and parse the token
	token, err := jwt.ParseWithClaims(session.Token, &Credential{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	// fmt.Println(session.Token[6:])
	// Extract the userID from the token claims
	_, ok := token.Claims.(*Credential)
	if !ok {
		return 0, errors.New("Failed to extract claims from token")
	}
	// db, err := con.GetDB()
	// if err != nil {
	// 	return 0, fmt.Errorf("failed to connect to the database: %v", err)
	// }

	var adminID int

	err = db.QueryRow("SELECT admin_id FROM admin_sessions WHERE token = ?", session.Token).Scan(&adminID)

	if err != nil {

		//fmt.Println(err)
		return 0, fmt.Errorf("failed to fetch adminID from the database: %v", err)
	}

	return adminID, nil
}
func LogOut(w http.ResponseWriter, r *http.Request) {

	// db, err := con.GetDB()
	// if err != nil {
	// 	panic(fmt.Errorf("failed to initialize database: %v", err))
	// }

	// Set the appropriate headers to enable CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent, Authorization")

	// Retrieve the JWT token from the request header
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {

		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&Response{Error: &Error{Message: "Unauthorized"}})
		return
	}

	// Extract the token from the Authorization header
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	session := &Session{
		Token: tokenString,
	}

	// Use the GetUserIDFromSession function to retrieve the userID
	sessionToken, err := GetSessionFromSession(session)

	if err != nil {
		logrus.Errorf("Error retrieving session token: %v", err)
		w.WriteHeader(http.StatusUnauthorized)
		//json.NewEncoder(w).Encode(&Response{Error: &Error{Message: "Unauthorized"}})
		return
	}

	// Delete the session from the database
	_, err = db.Exec("DELETE FROM admin_sessions WHERE token=?", sessionToken)
	if err != nil {
		logrus.Errorf("Error deleting session from the database: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return

	} else {
		response := Response{
			Data: "Successfully logged out",
		}

		jsonResponse, _ := json.Marshal(response)
		w.Write(jsonResponse)
		// w.Write("jdjkded")

	}

}
func GetSessionFromSession(session *Session) (string, error) {

	// Validate and parse the token
	token, err := jwt.ParseWithClaims(session.Token, &Credential{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	//fmt.Println(token)
	// Extract the userID from the token claims
	_, ok := token.Claims.(*Credential)
	if !ok {
		return "", errors.New("Failed to extract claims from token")
	}
	// db, err := con.GetDB()
	// if err != nil {
	// 	return "", fmt.Errorf("failed to connect to the database: %v", err)
	// }

	var session1 string

	err = db.QueryRow("SELECT token FROM admin_sessions WHERE token = ?", session.Token).Scan(&session1)

	if err != nil {

		//fmt.Println(err)
		return "", fmt.Errorf("failed to fetch session from the database: %v", err)
	}

	return session1, nil
}

var tmpl = template.Must(template.ParseGlob("form/*.html"))

func IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the token is present in the request's cookies
		cookie, err := r.Cookie("token")
		if err != nil || cookie == nil {
			tmpl.ExecuteTemplate(w, "admin.html", r)
			return
		}

		// Verify the token
		tokenString := cookie.Value

		session := &Session{
			Token: tokenString,
		}

		// Use the GetUserIDFromSession function to retrieve the userID
		sessionToken, err := GetSessionFromSession(session)

		ipAddress := getClientIP(r)

		// fmt.Println("+++++++++++++++++++++++++++++")
		// fmt.Println(ipAddress)

		ipDB, _ := GetIP(w, r)

		// fmt.Println("========================")

		// fmt.Println(ipDB)

		//tokenDB, err := GetSessionFromSession(session1)
		// claims := &Credential{}
		// token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// 	return jwtKey, nil // jwtKey is the same key used to sign the token in AdminLoginHandler
		// })

		if err != nil || sessionToken != tokenString || ipAddress != ipDB {
			tmpl.ExecuteTemplate(w, "admin.html", r)
			return
		}

		// If the token is valid, continue to the next handler
		next.ServeHTTP(w, r)

	})
}

func AdminDelete(w http.ResponseWriter, r *http.Request) {
	// _, err := con.GetDB()
	// if err != nil {
	// 	fmt.Errorf("failed to connect database", err)

	// }
	var admin AdminLogin
	params := mux.Vars(r)
	loanID := params["id"]

	result, err := db.Exec("DELETE FROM admin WHERE id=?", loanID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, "admin not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)

	logrus.WithFields(logrus.Fields{
		"admin_id": admin.ID,
	}).Warnln("Delete Successfully")
}

func getClientIP(r *http.Request) string {
	remoteAddr := r.RemoteAddr
	ip, _, err := net.SplitHostPort(remoteAddr)
	if err != nil {
		return remoteAddr
	}
	return ip
}

func GetIP(w http.ResponseWriter, r *http.Request) (string, error) {

	// db, _ = con.GetDB()

	authHeader := r.Header.Get("Cookie")

	//fmt.Println(authHeader)

	if authHeader == "" {
		logrus.Warnln("Authorization header missing")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&Response{Error: &Error{Message: "Unauthorized"}})
		return "", nil
	}

	// Extract the token from the Authorization header
	cookieParts := strings.Split(authHeader, "; ")
	var tokenString string
	for _, part := range cookieParts {
		if strings.HasPrefix(part, "token=") {
			tokenString = strings.TrimPrefix(part, "token=")
			break
		}
	}

	//fmt.Println(len(tokenString))
	var ip_address string

	err := db.QueryRow("SELECT ip_address FROM admin_sessions WHERE token = ?", tokenString).Scan(&ip_address)

	if err != nil {
		logrus.Infoln("failed to extract ip add", err)
	}

	return ip_address, nil
}

// Define a custom response writer to capture the output of template execution
type responseWriter struct {
	body []byte
}

func (rw *responseWriter) Header() http.Header {
	return make(http.Header)
}

func (rw *responseWriter) Write(data []byte) (int, error) {
	rw.body = data

	return len(data), nil
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	// Do nothing
}
func (rw *responseWriter) WriteTo(w http.ResponseWriter) {
	_, err := w.Write(rw.body)
	if err != nil {
		log.Fatal(err)
	}
}
func ProxyHandlerReact(w http.ResponseWriter, r *http.Request) {
	// Extract the URL from the request
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "URL parameter is missing", http.StatusBadRequest)
		return
	}

	// Create a new request to the external API
	req, err := http.NewRequest(http.MethodPost, url, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer req.Body.Close()

	// Set headers from the original request to the new request
	req.Header = r.Header

	// Make a request to the external API
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the appropriate headers to enable CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
	// Allow the following headers
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")

	// Handle preflight OPTIONS requests
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Copy the response headers to the actual response writer
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	// Write the response body to the actual response writer
	w.WriteHeader(resp.StatusCode)
	_, err = w.Write(body)
	if err != nil {
		//log.Println(err)
	}
}

func ProxyHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the URL from the request
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "URL parameter is missing", http.StatusBadRequest)
		return
	}

	// Create a new request to the external API
	req, err := http.NewRequest(r.Method, url, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer req.Body.Close()

	// Set headers from the original request to the new request
	req.Header = r.Header

	// Make a request to the external API
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the appropriate headers to enable CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
	// Allow the following headers
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")

	// Handle preflight OPTIONS requests
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Copy the response headers to the actual response writer
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	// Write the response body to the actual response writer
	w.WriteHeader(resp.StatusCode)
	_, err = w.Write(body)
	if err != nil {
		//log.Println(err)
	}
}
