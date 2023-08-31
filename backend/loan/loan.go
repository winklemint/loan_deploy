package loan

import (
	con "backend/Config"
	"backend/user"
	"database/sql"
	"encoding/json"
	"errors"
	"regexp"
	"strconv"
	"strings"

	"fmt"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"

	"net/http"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/mysql"

	_ "backend/user"

	"github.com/gorilla/mux"
)

type Loan_details struct {
	ID                   int       `json:"id,omitempty"`
	Full_Name            string    `json:"full_name,omitempty"`
	Applicant_Contact    string    `json:"applicant_contact,omitempty"`
	Loan_type            string    `json:"loan_type,omitempty" validate:"required"`
	Employment_type      string    `json:"employment_type,omitempty" `
	Loan_amount          float64   `json:"loan_amount,omitempty" validate:"gt=0"`
	Gross_monthly_income float64   `json:"gross_monthly_income"`
	Pincode              int       `json:"pincode,omitempty" validate:"len=6"`
	Tenure               int       `json:"tenure,omitempty" validate:"gt=0"`
	Created_at           time.Time `json:"created_at,omitempty"`
	Last_modified        time.Time `json:"last_modified,omitempty"`
	Status               string    `json:"status"`
	Remark               string    `json:"remark"`
	Admin_Name           string    `json:"admin_name"`
	Progress_Status      string    `json:"progress_status"`
	Tenure_flag          int       `json:"enq_moved_to_lead"`
	Device_ID            string    `json:"device_id"`
	jwt.StandardClaims
}

// type Credential struct {
// 	Name     string `json:"user_name"`
// 	Contact  string `json:"applicant_contact"`
// 	Loantype string `json:"loan_type"`
// 	jwt.StandardClaims
// }

var db *sql.DB
var loan Loan_details

func init() {
	var err error
	db, err = con.GetDB()
	if err != nil {
		panic(fmt.Errorf("failed to initialize database: %v", err))
	}

}

func ValidateLoan(loan Loan_details, validate *validator.Validate) error {
	// Check that loan_type is not empty
	if loan.Loan_type == "" {
		return errors.New("Loan type is required.")
	}

	// Check that loan_amount is greater than zero
	if loan.Loan_amount <= 0 {
		return errors.New("Loan amount must be greater than zero.")
	}

	// 	// Check that tenure is greater than zero
	//	if loan.Tenure <= 0 {
	// 		return errors.New("Tenure must be greater than zero.")
	// 	}

	// 	// Check that pincode is exactly 6 digits
	if match, _ := regexp.MatchString(`^\d{6}$`, fmt.Sprint(loan.Pincode)); !match {
		return errors.New("Pincode must be exactly 6 digits.")
	}

	// 	// Check that employment_type is not empty
	if loan.Employment_type == "" {
		return errors.New("Employment type is required.")
	}

	if loan.Gross_monthly_income <= 0 {
		return errors.New("Gross monthly income must be greater than zero.")
	}

	// 	// If all checks pass, return nil
	return nil
}

type Error struct {
	Message string `json:"message"`
}
type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error *Error      `json:"error,omitempty"`
}

// func InsertLoanDetails(w http.ResponseWriter, r *http.Request) {

// 	config, err := con.LoadConfig("Config/config.yaml")
// 	if err != nil {
// 		panic(err)
// 	}

// 	con.ConnectDB(config)
// 	defer con.CloseDB()

// 	//r.Header.Get("User-Agent")
// 	//r.Header("User-Agent")

// 	var loan Loan_details
// 	if err := json.NewDecoder(r.Body).Decode(&loan); err != nil {
// 		logrus.Errorf("Error decoding request body: %v", err)
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode(&Response{Error: &Error{Message: err.Error()}})
// 		return
// 	}

// 	//create validator instance
// 	validate := validator.New()

// 	// validate loan details
// 	if err := ValidateLoan(loan, validate); err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode(&Response{Error: &Error{Message: err.Error()}})
// 		return
// 	}

// 	result, err := db.Exec("INSERT INTO loan_details_table(loan_type,loan_amount,pincode,tenure,employment_type,gross_monthly_income,created_at,last_modified) VALUES(?,?,?,?,?,?,NOW(),NOW())", loan.Loan_type, loan.Loan_amount, loan.Pincode, loan.Tenure, loan.Employment_type, loan.Gross_monthly_income)
// 	if err != nil {
// 		logrus.Errorf("Error executing SQL query: %v", err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(w).Encode(&Response{Error: &Error{Message: err.Error()}})
// 		return
// 	}

// 	id, _ := result.LastInsertId()
// 	loan.ID = int(id)
// 	logrus.WithFields(logrus.Fields{
// 		"loan_id":              loan.ID,
// 		"Loan_Type":            loan.Loan_type,
// 		"Loan_Amount":          loan.Loan_amount,
// 		"Tenure":               loan.Tenure,
// 		"Pincode":              loan.Pincode,
// 		"Employment_Type":      loan.Employment_type,
// 		"Gross_Monthly_Income": loan.Gross_monthly_income,
// 	}).Info("Loan Details Inserted")

// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(&Response{Data: loan})
// }

// func InsertLoanDetails(w http.ResponseWriter, r *http.Request) {

// 	// Retrieve the JWT token from the request header
// 	// authHeader := r.Header.Get("Cookie")
// 	// if authHeader == "token" {

// 	// 	logrus.Warnln("Authorization header missing")
// 	// 	w.WriteHeader(http.StatusUnauthorized)
// 	// 	json.NewEncoder(w).Encode(&Response{Error: &Error{Message: "Unauthorized"}})
// 	// 	return
// 	// }

// 	db, err := con.GetDB()
// 	if err != nil {
// 		panic(fmt.Errorf("failed to initialize database: %v", err))
// 	}

// 	// Set the appropriate headers to enable CORS
// 	// w.Header().Set("Access-Control-Allow-Origin", "*")
// 	// w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
// 	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent, Authorization")

// 	// // Handle preflight OPTIONS requests
// 	// if r.Method == "OPTIONS" {
// 	// 	w.WriteHeader(http.StatusOK)
// 	// 	return
// 	// }

// 	// // Check request method
// 	// if r.Method != http.MethodPost {
// 	// 	//http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 	// 	return
// 	// }

// 	// Extract the token from the Authorization header
// 	// tokenString := strings.TrimPrefix(authHeader, "token=")

// 	// fmt.Println(tokenString)

// 	// session := &user.Session{
// 	// 	Token: tokenString,
// 	// }

// 	// Use the GetUserIDFromSession function from the login package to retrieve the userID
// 	// userID, err := user.GetUserIDFromSession(session)
// 	// fmt.Println(userID)
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// 	logrus.Errorf("Error retrieving userID from session: %v", err)
// 	// 	w.WriteHeader(http.StatusUnauthorized)
// 	// 	json.NewEncoder(w).Encode(&Response{Error: &Error{Message: "Unauthorized"}})
// 	// 	return
// 	// }

// 	// Create validator instance
// 	//validate := validator.New()

// 	// Validate loan details
// 	// if err := ValidateLoan(loan, validate); err != nil {
// 	// 	w.WriteHeader(http.StatusBadRequest)
// 	// 	json.NewEncoder(w).Encode(&Response{Error: &Error{Message: err.Error()}})
// 	// 	return
// 	// }

// 	// if err != nil {
// 	// 	logrus.Errorf("Error connecting to the database: %v", err)
// 	// 	w.WriteHeader(http.StatusInternalServerError)
// 	// 	json.NewEncoder(w).Encode(&Response{Error: &Error{Message: "Internal Server Error"}})
// 	// 	return
// 	// }
// 	var loan Loan_details
// 	if err := json.NewDecoder(r.Body).Decode(&loan); err != nil {
// 		logrus.Infoln("Error decoding request body:", err)

// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode(&Response{Error: &Error{Message: err.Error()}})
// 		return
// 	}
// 	var credential Credential
// 	if err := json.NewDecoder(r.Body).Decode(&credential); err != nil {
// 		fmt.Println(err)
// 		http.Error(w, "Invalid request payload", http.StatusBadRequest)
// 		return
// 	}
// 	fmt.Println(credential.Applicant_Contact)
// 	loan_type, err := GetDetailFromDatabase(credential.Applicant_Contact)
// 	fmt.Println(loan_type)
// 	if loan_type != credential.Loan_type || credential.Loan_type == "" {

// 		result, err := db.Exec("INSERT INTO loan_details_table (full_name,applicant_contact,loan_type, employment_type, loan_amount, gross_monthly_income, pincode, tenure,   created_at, last_modified) VALUES (?,?,?, ?, ?, ?, ?, ?, NOW(), NOW())", loan.Full_Name, loan.Applicant_Contact, loan.Loan_type, loan.Employment_type, loan.Loan_amount, loan.Gross_monthly_income, loan.Pincode, loan.Tenure)
// 		if err != nil {
// 			logrus.Errorf("Error executing SQL query: %v", err)
// 			w.WriteHeader(http.StatusInternalServerError)
// 			json.NewEncoder(w).Encode(&Response{Error: &Error{Message: "Internal Server Error"}})
// 			return
// 		}

// 		id, _ := result.LastInsertId()
// 		loan.ID = int(id)

// 		loanID := strconv.Itoa(loan.ID)

// 		http.SetCookie(w, &http.Cookie{
// 			Name:  "loan_id",
// 			Value: loanID,
// 		})

// 		logrus.WithFields(logrus.Fields{
// 			"loan_id":              loan.ID,
// 			"Loan_Type":            loan.Loan_type,
// 			"Loan_Amount":          loan.Loan_amount,
// 			"Tenure":               loan.Tenure,
// 			"Pincode":              loan.Pincode,
// 			"Employment_Type":      loan.Employment_type,
// 			"Gross_Monthly_Income": loan.Gross_monthly_income,
// 		}).Info("Loan Details Inserted")

// 		w.WriteHeader(http.StatusCreated)
// 		json.NewEncoder(w).Encode(&Response{Data: loan})
// 	} else {

// 		fmt.Println("user is allready login")
// 	}
// }

func InsertLoanDetails(w http.ResponseWriter, r *http.Request) {
	// db, err := con.GetDB()
	// if err != nil {
	// 	panic(fmt.Errorf("failed to initialize database: %v", err))
	// }

	cookie, _ := r.Cookie("device_id")

	device := cookie.Value

	cookieName := "token"

	cookies, _ := r.Cookie(cookieName)

	token := cookies.Value

	if device != "" && token != "" {

		fmt.Println("+++++++++++++++++++++++")

		tokenString := strings.TrimPrefix(token, "token=")

		session := &user.Session{
			Token: tokenString,
		}

		//Use the GetUserIDFromSession function from the login package to retrieve the userID
		userID, err := user.GetUserIDFromSession(session)
		fmt.Println(userID)
		if err != nil {
			fmt.Println(err)
			logrus.Errorf("Error retrieving userID from session: %v", err)
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(&Response{Error: &Error{Message: "Unauthorized"}})
			return
		}

		var loan Loan_details
		if err := json.NewDecoder(r.Body).Decode(&loan); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		fmt.Println(loan.Applicant_Contact)

		// Check if the applicant_contact already exists in the database
		loan_type, _ := GetDetailFromDatabase(loan.Applicant_Contact, loan.Loan_type)
		if loan_type == "" {
			db, err := con.GetDB()
			if err != nil {
				panic(fmt.Errorf("failed to initialize database: %v", err))
			}

			result, err := db.Exec("INSERT INTO loan_details_table (full_name,applicant_contact,loan_type, employment_type, loan_amount, gross_monthly_income, pincode, tenure, user_id) VALUES (?,?,?, ?, ?, ?, ?, ?, ?)", loan.Full_Name, loan.Applicant_Contact, loan.Loan_type, loan.Employment_type, loan.Loan_amount, loan.Gross_monthly_income, loan.Pincode, loan.Tenure, userID)
			if err != nil {
				logrus.Infoln(err)
				http.Error(w, "Database error", http.StatusInternalServerError)
				return
			}

			id, _ := result.LastInsertId()
			loan.ID = int(id)

			loanID := strconv.Itoa(loan.ID)

			http.SetCookie(w, &http.Cookie{
				Name:  "loan_id",
				Value: loanID,
			})

			logrus.WithFields(logrus.Fields{
				"loan_id":              loan.ID,
				"Loan_Type":            loan.Loan_type,
				"Loan_Amount":          loan.Loan_amount,
				"Tenure":               loan.Tenure,
				"Pincode":              loan.Pincode,
				"Employment_Type":      loan.Employment_type,
				"Gross_Monthly_Income": loan.Gross_monthly_income,
			}).Info("Loan Details Inserted")

			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(&Response{Data: result})

			return
		}

		if loan_type == loan.Loan_type {
			// User with the same applicant_contact already has a loan with the same loan_type
			fmt.Println("User has already added a loan with this loan_type")
			http.Error(w, "User has already added a loan with this loan_type", http.StatusConflict)
			return
		} else {

			// User can proceed to add a new loan since the loan_type is different
			result, err := db.Exec("INSERT INTO loan_details_table (full_name,applicant_contact,loan_type, employment_type, loan_amount, gross_monthly_income, pincode, tenure, user_id) VALUES (?,?,?, ?, ?, ?, ?, ?, ?)", loan.Full_Name, loan.Applicant_Contact, loan.Loan_type, loan.Employment_type, loan.Loan_amount, loan.Gross_monthly_income, loan.Pincode, loan.Tenure, userID)
			if err != nil {
				http.Error(w, "Database error", http.StatusInternalServerError)
				return
			}

			id, _ := result.LastInsertId()
			loan.ID = int(id)

			loanID := strconv.Itoa(loan.ID)

			http.SetCookie(w, &http.Cookie{
				Name:  "loan_id",
				Value: loanID,
			})

			logrus.WithFields(logrus.Fields{
				"loan_id":              loan.ID,
				"Loan_Type":            loan.Loan_type,
				"Loan_Amount":          loan.Loan_amount,
				"Tenure":               loan.Tenure,
				"Pincode":              loan.Pincode,
				"Employment_Type":      loan.Employment_type,
				"Gross_Monthly_Income": loan.Gross_monthly_income,
			}).Info("Loan Details Inserted")

			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(&Response{Data: result})
		}

	} else if device != "" {

		fmt.Println("#######################")

		deviceString := strings.TrimPrefix(device, "device_id=")

		var loan Loan_details
		if err := json.NewDecoder(r.Body).Decode(&loan); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		fmt.Println(loan.Applicant_Contact)

		// Check if the applicant_contact already exists in the database
		loan_type, _ := GetDetailFromDatabase(loan.Applicant_Contact, loan.Loan_type)
		if loan_type == "" {
			// db, err := con.GetDB()
			// if err != nil {
			// 	panic(fmt.Errorf("failed to initialize database: %v", err))
			// }

			result, err := db.Exec("INSERT INTO loan_details_table (full_name,applicant_contact,loan_type, employment_type, loan_amount, gross_monthly_income, pincode, tenure, device_id) VALUES (?,?,?, ?, ?, ?, ?, ?, ?)", loan.Full_Name, loan.Applicant_Contact, loan.Loan_type, loan.Employment_type, loan.Loan_amount, loan.Gross_monthly_income, loan.Pincode, loan.Tenure, deviceString)
			if err != nil {
				logrus.Infoln(err)
				http.Error(w, "Database error", http.StatusInternalServerError)
				return
			}

			id, _ := result.LastInsertId()
			loan.ID = int(id)

			loanID := strconv.Itoa(loan.ID)

			http.SetCookie(w, &http.Cookie{
				Name:  "loan_id",
				Value: loanID,
			})

			logrus.WithFields(logrus.Fields{
				"loan_id":              loan.ID,
				"Loan_Type":            loan.Loan_type,
				"Loan_Amount":          loan.Loan_amount,
				"Tenure":               loan.Tenure,
				"Pincode":              loan.Pincode,
				"Employment_Type":      loan.Employment_type,
				"Gross_Monthly_Income": loan.Gross_monthly_income,
			}).Info("Loan Details Inserted")

			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(&Response{Data: result})
			return
		}

		if loan_type == loan.Loan_type {
			// User with the same applicant_contact already has a loan with the same loan_type
			fmt.Println("User has already added a loan with this loan_type")
			http.Error(w, "User has already added a loan with this loan_type", http.StatusConflict)
			return
		} else {

			// User can proceed to add a new loan since the loan_type is different
			result, err := db.Exec("INSERT INTO loan_details_table (full_name,applicant_contact,loan_type, employment_type, loan_amount, gross_monthly_income, pincode, tenure, device_id) VALUES (?,?,?, ?, ?, ?, ?, ?, ?)", loan.Full_Name, loan.Applicant_Contact, loan.Loan_type, loan.Employment_type, loan.Loan_amount, loan.Gross_monthly_income, loan.Pincode, loan.Tenure, deviceString)
			if err != nil {
				http.Error(w, "Database error", http.StatusInternalServerError)
				return
			}

			id, _ := result.LastInsertId()
			loan.ID = int(id)

			loanID := strconv.Itoa(loan.ID)

			http.SetCookie(w, &http.Cookie{
				Name:  "loan_id",
				Value: loanID,
			})

			logrus.WithFields(logrus.Fields{
				"loan_id":              loan.ID,
				"Loan_Type":            loan.Loan_type,
				"Loan_Amount":          loan.Loan_amount,
				"Tenure":               loan.Tenure,
				"Pincode":              loan.Pincode,
				"Employment_Type":      loan.Employment_type,
				"Gross_Monthly_Income": loan.Gross_monthly_income,
			}).Info("Loan Details Inserted")

			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(&Response{Data: result})
		}

	} else {
		result, err := db.Exec("INSERT INTO loan_details_table (full_name,applicant_contact,loan_type, employment_type, loan_amount, gross_monthly_income, pincode, tenure) VALUES (?,?,?, ?, ?, ?, ?, ?, ?)", loan.Full_Name, loan.Applicant_Contact, loan.Loan_type, loan.Employment_type, loan.Loan_amount, loan.Gross_monthly_income, loan.Pincode, loan.Tenure)
		if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}

		id, _ := result.LastInsertId()
		loan.ID = int(id)

		loanID := strconv.Itoa(loan.ID)

		http.SetCookie(w, &http.Cookie{
			Name:  "loan_id",
			Value: loanID,
		})

		logrus.WithFields(logrus.Fields{
			"loan_id":              loan.ID,
			"Loan_Type":            loan.Loan_type,
			"Loan_Amount":          loan.Loan_amount,
			"Tenure":               loan.Tenure,
			"Pincode":              loan.Pincode,
			"Employment_Type":      loan.Employment_type,
			"Gross_Monthly_Income": loan.Gross_monthly_income,
		}).Info("Loan Details Inserted")

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(&Response{Data: result})
	}

}

// deviceString := strings.TrimPrefix(device, "device_id=")

// 		var loan Loan_details
// 		if err := json.NewDecoder(r.Body).Decode(&loan); err != nil {
// 			http.Error(w, "Invalid request payload", http.StatusBadRequest)
// 			return
// 		}

// 		fmt.Println(loan.Applicant_Contact)

// 		// Check if the applicant_contact already exists in the database
// 		loan_type, _ := GetDetailFromDatabase(loan.Applicant_Contact, loan.Loan_type)
// 		if loan_type == "" {
// 			db, err := con.GetDB()
// 			if err != nil {
// 				panic(fmt.Errorf("failed to initialize database: %v", err))
// 			}

// 			result, err := db.Exec("INSERT INTO loan_details_table (full_name,applicant_contact,loan_type, employment_type, loan_amount, gross_monthly_income, pincode, tenure, device_id) VALUES (?,?,?, ?, ?, ?, ?, ?, ?)", loan.Full_Name, loan.Applicant_Contact, loan.Loan_type, loan.Employment_type, loan.Loan_amount, loan.Gross_monthly_income, loan.Pincode, loan.Tenure, deviceString)
// 			if err != nil {
// 				logrus.Infoln(err)
// 				http.Error(w, "Database error", http.StatusInternalServerError)
// 				return
// 			}

// 			w.WriteHeader(http.StatusCreated)
// 			json.NewEncoder(w).Encode(&Response{Data: result})
// 			return
// 		}

// 		if loan_type == loan.Loan_type {
// 			// User with the same applicant_contact already has a loan with the same loan_type
// 			fmt.Println("User has already added a loan with this loan_type")
// 			http.Error(w, "User has already added a loan with this loan_type", http.StatusConflict)
// 			return
// 		} else {

// 			// User can proceed to add a new loan since the loan_type is different
// 			result, err := db.Exec("INSERT INTO loan_details_table (full_name,applicant_contact,loan_type, employment_type, loan_amount, gross_monthly_income, pincode, tenure, device_id) VALUES (?,?,?, ?, ?, ?, ?, ?, ?)", loan.Full_Name, loan.Applicant_Contact, loan.Loan_type, loan.Employment_type, loan.Loan_amount, loan.Gross_monthly_income, loan.Pincode, loan.Tenure, deviceString)
// 			if err != nil {
// 				http.Error(w, "Database error", http.StatusInternalServerError)
// 				return
// 			}

// 			w.WriteHeader(http.StatusCreated)
// 			json.NewEncoder(w).Encode(&Response{Data: result})
// 		}

func GetDetailFromDatabase(contact, loan_type string) (string, error) {
	fmt.Println("Get details from database", contact)
	// db, _ := con.GetDB()
	loan := "null"
	fmt.Println(loan)
	err := db.QueryRow("SELECT loan_type FROM loan_details_table WHERE applicant_contact = ? and loan_type=?", contact, loan_type).Scan(&loan)
	if err != nil {
		fmt.Println(err)

		return "", err
	}
	fmt.Println(loan)
	return loan, err

}

// for admin panel

func UpdateLoanDetailsForAdmin(w http.ResponseWriter, r *http.Request) {
	// db, err := con.GetDB()
	// if err != nil {
	// 	panic(fmt.Errorf("failed to initialize database: %v", err))
	// }
	params := mux.Vars(r)
	LoanID := params["id"]
	var loan Loan_details
	//fmt.Println("r.body", r.Body)
	error := json.NewDecoder(r.Body).Decode(&loan)
	if error != nil {
		http.Error(w, error.Error(), http.StatusBadRequest)
		return
	}
	result, err := db.Exec("UPDATE loan_details_table SET loan_type=?, loan_amount=?, tenure=?, Pincode=?, employment_type=?, gross_monthly_income=?,status=?,remark=?,admin_name=? WHERE id=?", loan.Loan_type, loan.Loan_amount, loan.Tenure, loan.Pincode, loan.Employment_type, loan.Gross_monthly_income, loan.Status, loan.Remark, loan.Admin_Name, LoanID)
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
		http.Error(w, "Loan ID not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Loan ID updated successfully")
	logrus.WithFields(logrus.Fields{
		"loan_id":              LoanID,
		"Loan_Type":            loan.Loan_type,
		"Loan_Amount":          loan.Loan_amount,
		"Tenure":               loan.Tenure,
		"Pincode":              loan.Pincode,
		"Employment_Type":      loan.Employment_type,
		"Gross_Monthly_Income": loan.Gross_monthly_income,
	}).Info("Update Loan Table Successfully")
	// err = templates.ExecuteTemplate(w, "http://localhost:9000/form/update_loan.html", loan)
	// if err != nil {
	//  http.Error(w, err.Error(), http.StatusInternalServerError)
	//  return
	// }
}

func UpdateLoanDetails(w http.ResponseWriter, r *http.Request) {
	// db, err := con.GetDB()
	// if err != nil {
	// 	panic(fmt.Errorf("failed to initialize database: %v", err))
	// }

	cookie, _ := r.Cookie("device_id")

	device := cookie.Value

	cookieName := "token"

	cookies, _ := r.Cookie(cookieName)

	token := cookies.Value

	if device != "" && token != "" {

		tokenString := strings.TrimPrefix(token, "token=")

		session := &user.Session{
			Token: tokenString,
		}

		//Use the GetUserIDFromSession function from the login package to retrieve the userID
		userID, err := user.GetUserIDFromSession(session)
		fmt.Println(userID)
		if err != nil {
			fmt.Println(err)
			logrus.Errorf("Error retrieving userID from session: %v", err)
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(&Response{Error: &Error{Message: "Unauthorized"}})
			return
		}

		params := mux.Vars(r)
		LoanID := params["id"]

		var loan Loan_details
		error := json.NewDecoder(r.Body).Decode(&loan)
		if error != nil {
			http.Error(w, error.Error(), http.StatusBadRequest)
			return
		}

		// Initialize the query and arguments
		query := "UPDATE loan_details_table SET"
		args := []interface{}{}

		query += "user_id=?"
		args = append(args, userID)

		if loan.Loan_type != "" {
			query += " loan_type=?"
			args = append(args, loan.Loan_type)
		}
		if loan.Loan_amount > 0 {
			query += " loan_amount=?"
			args = append(args, loan.Loan_amount)
		}
		if loan.Tenure > 0 {
			query += " tenure=? , enq_moved_to_lead=1 "
			args = append(args, loan.Tenure)
		}

		if loan.Pincode >= 6 {
			query += " pincode=?"
			args = append(args, loan.Pincode)
		}

		if loan.Employment_type != "" {
			query += " employment_type=?"
			args = append(args, loan.Employment_type)
		}

		if loan.Gross_monthly_income > 0 {
			query += " gross_monthly_income=?"
			args = append(args, loan.Gross_monthly_income)
		}

		// if len(args) > 0 {
		// 	query = query[:len(query)-1]
		// }

		// Add the WHERE clause and the LoanID argument
		query += " WHERE id=? "
		args = append(args, LoanID)

		// Execute the query
		result, err := db.Exec(query, args...)
		if err != nil {
			logrus.Errorf("Error updating sql query for loan %v", err)
		}

		err = db.QueryRow("SELECT enq_moved_to_lead FROM loan_details_table WHERE id =?", LoanID).Scan(&loan.Tenure_flag)

		if loan.Tenure_flag == 1 {
			_, err := db.Exec("INSERT INTO lead_table SELECT * FROM loan_details_table  WHERE id = ?",
				LoanID)
			if err != nil {
				logrus.Errorf("Error inserting sql query for loan to lead %v", err)
			}

			_, err = db.Exec("UPDATE loan_details_table SET is_delete = 1 WHERE id =?", LoanID)
			if err != nil {
				logrus.Errorf("Error executing SQL query for update is_delete: %v", err)
			}

		}

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
			http.Error(w, "Loan ID not found", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Loan ID updated successfully")
		logrus.WithFields(logrus.Fields{
			"loan_id":              LoanID,
			"Loan_Type":            loan.Loan_type,
			"Loan_Amount":          loan.Loan_amount,
			"Tenure":               loan.Tenure,
			"Pincode":              loan.Pincode,
			"Employment_Type":      loan.Employment_type,
			"Gross_Monthly_Income": loan.Gross_monthly_income,
		}).Info("Update Loan Table Successfully")

		// err = templates.ExecuteTemplate(w, "http://localhost:9000/form/update_loan.html", loan)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }

	} else {

		params := mux.Vars(r)
		LoanID := params["id"]

		var loan Loan_details
		error := json.NewDecoder(r.Body).Decode(&loan)
		if error != nil {
			http.Error(w, error.Error(), http.StatusBadRequest)
			return
		}

		// Initialize the query and arguments
		query := "UPDATE loan_details_table SET"
		args := []interface{}{}

		if loan.Loan_type != "" {
			query += " loan_type=?"
			args = append(args, loan.Loan_type)
		}
		if loan.Loan_amount > 0 {
			query += " loan_amount=?"
			args = append(args, loan.Loan_amount)
		}
		if loan.Tenure > 0 {
			query += " tenure=? , enq_moved_to_lead=1 "
			args = append(args, loan.Tenure)
		}

		if loan.Pincode >= 6 {
			query += " pincode=?"
			args = append(args, loan.Pincode)
		}

		if loan.Employment_type != "" {
			query += " employment_type=?"
			args = append(args, loan.Employment_type)
		}

		if loan.Gross_monthly_income > 0 {
			query += " gross_monthly_income=?"
			args = append(args, loan.Gross_monthly_income)
		}

		// if len(args) > 0 {
		// 	query = query[:len(query)-1]
		// }

		// Add the WHERE clause and the LoanID argument
		query += " WHERE id=? "
		args = append(args, LoanID)

		// Execute the query
		result, err := db.Exec(query, args...)

		err = db.QueryRow("SELECT enq_moved_to_lead FROM loan_details_table WHERE id =?", LoanID).Scan(&loan.Tenure_flag)

		if loan.Tenure_flag == 1 {
			_, err := db.Exec("INSERT INTO lead_table SELECT * FROM loan_details_table  WHERE id = ?",
				LoanID)

			_, err = db.Exec("UPDATE loan_details_table SET is_delete = 1 WHERE id =?", LoanID)
			if err != nil {
				logrus.Errorf("Error executing SQL query: %v", err)
			}

		}

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
			http.Error(w, "Loan ID not found", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Loan ID updated successfully")
		logrus.WithFields(logrus.Fields{
			"loan_id":              LoanID,
			"Loan_Type":            loan.Loan_type,
			"Loan_Amount":          loan.Loan_amount,
			"Tenure":               loan.Tenure,
			"Pincode":              loan.Pincode,
			"Employment_Type":      loan.Employment_type,
			"Gross_Monthly_Income": loan.Gross_monthly_income,
		}).Info("Update Loan Table Successfully")

		// err = templates.ExecuteTemplate(w, "http://localhost:9000/form/update_loan.html", loan)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }
	}
}

func DeleteLoanDetails(w http.ResponseWriter, r *http.Request) {

	// db, err := con.GetDB()
	// if err != nil {
	// 	panic(fmt.Errorf("failed to initialize database: %v", err))
	// }

	params := mux.Vars(r)
	loanID := params["id"]

	result, err := db.Exec("DELETE FROM loan_details_table WHERE id=?", loanID)
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
		http.Error(w, "Loan not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "loan deleted successfully")
	logrus.WithFields(logrus.Fields{
		"loan_id": loan.ID,
	}).Warnln("Delete Successfully")
}

//var templates = template.Must(template.ParseFiles("form/loan.html"))

func GetLoanDetails(w http.ResponseWriter, r *http.Request) {
	// db, err := con.GetDB()
	// if err != nil {
	// 	panic(fmt.Errorf("failed to initialize database: %v", err))
	// }

	logrus.Infoln("Connection Start")
	defer con.CloseDB()

	params := mux.Vars(r)
	loanID := params["id"]

	rows, err := db.Query("SELECT user_info.id, user_info.user_name, user_info.user_email, user_info.user_contact_num , loan_details_table.id,loan_details_table.loan_type,loan_details_table.loan_amount, loan_details_table.tenure, loan_details_table.pincode, loan_details_table.employment_type, loan_details_table.gross_monthly_income, loan_details_table.status, loan_details_table.created_at,loan_details_table.remark,admin.username FROM loan_details_table INNER JOIN user_info ON loan_details_table.user_id = user_info.id  JOIN admin ON loan_details_table.admin_name = admin.id WHERE loan_details_table.id=?", loanID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()
	var data []map[string]interface{}

	for rows.Next() {

		var user user.User_info
		var loan Loan_details
		var CreatedAtstr string
		err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Contact_Num, &loan.ID, &loan.Loan_type, &loan.Loan_amount, &loan.Tenure, &loan.Pincode, &loan.Employment_type, &loan.Gross_monthly_income, &loan.Status, &CreatedAtstr, &loan.Remark, &loan.Admin_Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		CreatedAt, err := time.Parse("2006-01-02 15:04:05", CreatedAtstr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		loan.Created_at = CreatedAt

		// Create a map for each lead and user, then append them to the data slice
		leadMap := map[string]interface{}{
			"id":                   loan.ID,
			"loan_type":            loan.Loan_type,
			"employment_type":      loan.Employment_type,
			"loan_amount":          loan.Loan_amount,
			"gross_monthly_income": loan.Gross_monthly_income,
			"pincode":              loan.Pincode,
			"tenure":               loan.Tenure,
			"status":               loan.Status,
			"created_at":           loan.Created_at,
			"user_id":              user.Id,
			"user_name":            user.Name,
			"user_email":           user.Email,
			"user_contact":         user.Contact_Num,
			"remark":               loan.Remark, // Include the 'remark' field in the leadMap
			"admin_name":           loan.Admin_Name,
		}

		data = append(data, leadMap)

	}

	jsonResponse, err := json.Marshal(Response{
		Data: data,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

//retrieve all data from loan

// func LoanIndex(w http.ResponseWriter, r *http.Request) {
// 	config, err := con.LoadConfig("Config/config.yaml")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	con.ConnectDB(config)
// 	logrus.Infoln("Connection Start")
// 	defer con.CloseDB()
// 	fmt.Println("=====================")

// 	rows, err := db.Query("SELECT user_info.id, user_info.user_name, user_info.user_email, user_info.user_contact_num ,loan_details_table.id, loan_details_table.loan_type, loan_details_table.loan_amount, loan_details_table.tenure, loan_details_table.pincode, loan_details_table.employment_type, loan_details_table.gross_monthly_income, loan_details_table.status, loan_details_table.created_at, loan_details_table.last_modified,loan_details_table.remark, admin.username FROM loan_details_table INNER JOIN user_info ON loan_details_table.user_id = user_info.id JOIN admin ON loan_details_table.admin_name = admin.id WHERE loan_details_table.is_delete = 0 ")

// 	fmt.Println("==============")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	//defer rows.Close()
// 	fmt.Println("=================")
// 	var data []map[string]interface{}
// 	for rows.Next() {
// 		var user user.User_info
// 		var loan Loan_details
// 		var CreatedAtstr, LastAtStr string
// 		fmt.Println("====================")

// 		err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Contact_Num, &loan.ID, &loan.Loan_type, &loan.Loan_amount, &loan.Tenure, &loan.Pincode, &loan.Employment_type, &loan.Gross_monthly_income, &loan.Status, &CreatedAtstr, &LastAtStr, &loan.Remark, &loan.Admin_Name)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		// loan.Status = getStatusValue(statusText)
// 		CreatedAt, err := time.Parse("2006-01-02 15:04:05", CreatedAtstr)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 		}
// 		loan.Created_at = CreatedAt
// 		LastAt, err := time.Parse("2006-01-02 15:04:05", LastAtStr)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 		}
// 		loan.Last_modified = LastAt

// 		loanMap := map[string]interface{}{
// 			"id":                   loan.ID,
// 			"loan_type":            loan.Loan_type,
// 			"employment_type":      loan.Employment_type,
// 			"loan_amount":          loan.Loan_amount,
// 			"gross_monthly_income": loan.Gross_monthly_income,
// 			"pincode":              loan.Pincode,
// 			"tenure":               loan.Tenure,
// 			"status":               loan.Status,
// 			"created_at":           loan.Created_at,
// 			"last_modified":        loan.Last_modified,
// 			"user_id":              user.Id,
// 			"user_name":            user.Name,
// 			"user_email":           user.Email,
// 			"user_contact":         user.Contact_Num,
// 			"remark":               loan.Remark, // Include the 'remark' field in the loanMap
// 			"admin_name":           loan.Admin_Name,
// 		}
// 		fmt.Println(data)
// 		fmt.Println(loanMap)

// 		data = append(data, loanMap)
// 		fmt.Println(data)

// 	}

// 	jsonResponse, err := json.Marshal(Response{
// 		Data: data,
// 	})
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(jsonResponse)
// }

func LoanIndex1(w http.ResponseWriter, r *http.Request) {
	// db, _ = con.GetDB()
	// defer db.Close()

	rows, err := db.Query("SELECT user_info.id, user_info.user_name, user_info.user_email, user_info.user_contact_num , loan_details_table.id,loan_details_table.loan_type, loan_details_table.loan_amount, loan_details_table.tenure, loan_details_table.pincode, loan_details_table.employment_type, loan_details_table.gross_monthly_income, loan_details_table.status, loan_details_table.created_at,loan_details_table.remark,admin.username,loan_details_table.progress_status FROM loan_details_table INNER JOIN user_info ON loan_details_table.user_id = user_info.id INNER JOIN admin ON loan_details_table.admin_name = admin.id ")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var data []map[string]interface{}

	for rows.Next() {

		var user user.User_info
		var loan Loan_details
		var CreatedAtstr string
		err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Contact_Num, &loan.ID, &loan.Loan_type, &loan.Loan_amount, &loan.Tenure, &loan.Pincode, &loan.Employment_type, &loan.Gross_monthly_income, &loan.Status, &CreatedAtstr, &loan.Remark, &loan.Admin_Name, &loan.Progress_Status)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		CreatedAt, err := time.Parse("2006-01-02 15:04:05", CreatedAtstr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		loan.Created_at = CreatedAt

		// Create a map for each lead and user, then append them to the data slice
		leadMap := map[string]interface{}{
			"id":                   loan.ID,
			"loan_type":            loan.Loan_type,
			"employment_type":      loan.Employment_type,
			"loan_amount":          loan.Loan_amount,
			"gross_monthly_income": loan.Gross_monthly_income,
			"pincode":              loan.Pincode,
			"tenure":               loan.Tenure,
			"status":               loan.Status,
			"created_at":           loan.Created_at,
			"user_id":              user.Id,
			"user_name":            user.Name,
			"user_email":           user.Email,
			"user_contact":         user.Contact_Num,
			"remark":               loan.Remark,
			"admin_name":           loan.Admin_Name,
			"progress_status":      loan.Progress_Status,
		}

		data = append(data, leadMap)

	}

	jsonResponse, err := json.Marshal(Response{
		Data: data,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
func LoanIndex(w http.ResponseWriter, r *http.Request) {
	// db, err := con.GetDB()
	// if err != nil {
	// 	panic(fmt.Errorf("failed to initialize database: %v", err))
	// }
	// defer db.Close()
	params := mux.Vars(r)
	pagesStr := params["page"]

	// Convert the pagesStr to an integer (page number)
	pages, err := strconv.Atoi(pagesStr)
	if err != nil {
		http.Error(w, "Invalid page number", http.StatusBadRequest)
		return
	}

	pageSize := 4
	offset := (pages - 1) * pageSize
	if err != nil {
		panic(fmt.Errorf("failed to initialize database: %v", err))
	}
	rows, err := db.Query("SELECT user_info.id, user_info.user_name, user_info.user_email, user_info.user_contact_num , loan_details_table.id,loan_details_table.loan_type, loan_details_table.loan_amount, loan_details_table.tenure, loan_details_table.pincode, loan_details_table.employment_type, loan_details_table.gross_monthly_income, loan_details_table.status, loan_details_table.created_at,loan_details_table.remark,admin.username,loan_details_table.progress_status FROM loan_details_table INNER JOIN user_info ON loan_details_table.user_id = user_info.id INNER JOIN admin ON loan_details_table.admin_name = admin.id WHERE loan_details_table.is_delete=0  order by loan_details_table.id desc limit 4 offset ?", offset)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var data []map[string]interface{}

	for rows.Next() {

		var user user.User_info
		var loan Loan_details
		var CreatedAtstr string
		err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Contact_Num, &loan.ID, &loan.Loan_type, &loan.Loan_amount, &loan.Tenure, &loan.Pincode, &loan.Employment_type, &loan.Gross_monthly_income, &loan.Status, &CreatedAtstr, &loan.Remark, &loan.Admin_Name, &loan.Progress_Status)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		CreatedAt, err := time.Parse("2006-01-02 15:04:05", CreatedAtstr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		loan.Created_at = CreatedAt

		// Create a map for each lead and user, then append them to the data slice
		leadMap := map[string]interface{}{
			"id":                   loan.ID,
			"loan_type":            loan.Loan_type,
			"employment_type":      loan.Employment_type,
			"loan_amount":          loan.Loan_amount,
			"gross_monthly_income": loan.Gross_monthly_income,
			"pincode":              loan.Pincode,
			"tenure":               loan.Tenure,
			"status":               loan.Status,
			"created_at":           loan.Created_at,
			"user_id":              user.Id,
			"user_name":            user.Name,
			"user_email":           user.Email,
			"user_contact":         user.Contact_Num,
			"remark":               loan.Remark,
			"admin_name":           loan.Admin_Name,
			"progress_status":      loan.Progress_Status,
		}

		data = append(data, leadMap)

	}

	jsonResponse, err := json.Marshal(Response{
		Data: data,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
func UpdateProgress(w http.ResponseWriter, r *http.Request) {
	// db, err := con.GetDB()
	// if err != nil {
	// 	panic(fmt.Errorf("failed to initialize database: %v", err))
	// }
	// defer db.Close()

	// Set the appropriate headers to enable CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent, Authorization")

	// Handle preflight OPTIONS requests
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Check request method
	if r.Method != http.MethodPatch {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	params := mux.Vars(r)
	LoanID := params["id"]
	var loan Loan_details

	err := json.NewDecoder(r.Body).Decode(&loan)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//fmt.Println(LeadID)
	result, err := db.Exec("UPDATE loan_details_table SET progress_status=? WHERE id=?", loan.Progress_Status, LoanID)

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
		http.Error(w, "Lead ID not found", http.StatusNotFound)
		return
	} else {
		fmt.Println("lead update succesfully")
	}

	//w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Lead progress updated successfully")
	logrus.WithFields(logrus.Fields{
		"Progress Status": loan.Progress_Status,
	}).Info("Update Lead Table Successfully")
}
func LoanCount(w http.ResponseWriter, r *http.Request) {
	// db, err := con.GetDB()
	// if err != nil {
	// 	panic(fmt.Errorf("failed to initialize database: %v", err))
	// }

	//db := con.GetDB()

	rows, err := db.Query("SELECT COUNT(*) as counting FROM loan_details_table WHERE is_delete=0")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var counting int

	for rows.Next() {
		err = rows.Scan(&counting)
		if err != nil {
			panic(err)
		}
	}

	// loans := make([]loan_info, 0)
	// loan := loan_info{} // Assuming loan_info is a struct type

	// loans = append(loans, loan)

	// response := Response{
	// 	Data: loans,
	// }

	jsonResponse, err := json.Marshal(counting)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

//admin soft delete

func AdminSoftDeleteLoan(w http.ResponseWriter, r *http.Request) {
	// db, err := con.GetDB()
	// if err != nil {
	// 	panic(fmt.Errorf("failed to initialize database: %v", err))
	// }

	params := mux.Vars(r)
	LeadID := params["id"]

	if r.Method != http.MethodPatch {
		w.Header().Set("Allow", http.MethodPatch)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Update is_delete field
	result, err := db.Exec("UPDATE loan_details_table SET is_delete=1 WHERE id=?", LeadID)
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
		http.Error(w, "Lead ID not found", http.StatusNotFound)
		return
	}
	// w.WriteHeader(http.StatusOK)
	// fmt.Fprintf(w, "Loan ID updated successfully")

}

// func LoanIndex1(w http.ResponseWriter, r *http.Request) {
// 	db, err := con.GetDB()
// 	if err != nil {
// 		panic(fmt.Errorf("failed to initialize database: %v", err))
// 	}

// 	rows, err := db.Query("SELECT id, loan_type, loan_amount, tenure, pincode, employment_type, gross_monthly_income, status, created_at, last_modified FROM loan_details_table")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	defer rows.Close()

// 	var loans []Loan_details

// 	for rows.Next() {
// 		var loan Loan_details
// 		var CreatedAtstr, LastAtStr string

// 		err = rows.Scan(&loan.ID, &loan.Loan_type, &loan.Loan_amount, &loan.Tenure, &loan.Pincode, &loan.Employment_type, &loan.Gross_monthly_income, &loan.Status, &CreatedAtstr, &LastAtStr)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		// loan.Status = getStatusValue(statusText)
// 		CreatedAt, err := time.Parse("2006-01-02 15:04:05", CreatedAtstr)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 		}
// 		loan.Created_at = CreatedAt
// 		LastAt, err := time.Parse("2006-01-02 15:04:05", LastAtStr)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 		}
// 		loan.Last_modified = LastAt

// 		loans = append(loans, loan)
// 	}

// 	response := Response{
// 		Data: loans,
// 	}

// 	jsonResponse, err := json.Marshal(response)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(jsonResponse)

// }
