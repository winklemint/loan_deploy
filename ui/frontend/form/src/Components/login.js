import React, { useState } from 'react';
import { useNavigate, Link } from 'react-router-dom';

const Login = ({ toggle }) => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [contact_num, setContactNum] = useState('');
  const [errors, setErrors] = useState({});
  const navigate = useNavigate();

  const handleSubmit = (e) => {
    e.preventDefault();

    let formErrors = {};
    if (!email) {
      formErrors.email = 'Please enter your email';
    } else if (!validateEmail(email)) {
      formErrors.email = 'Please enter a valid email';
    }
    if (!contact_num) {
      formErrors.contact_num = 'Please enter your phone contact number';
    } else if (!validatePhone(contact_num)) {
      formErrors.contact_num = 'Please enter a valid phone number';
    }
    if (!password) {
      formErrors.password = 'Please enter your password';
    } else if (password.length < 8) {
      formErrors.password = 'Password must contain at least 8 characters';
    }
    setErrors(formErrors);

    if (
      Object.keys(formErrors).length === 0 &&
      email &&
      contact_num &&
      password
    ) {
      fetch('/proxy?url=http://127.0.0.1:8080/user/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          email: email,
          contact_num: parseInt(contact_num),
          password: password,
        }),
      })
        .then((response) => {
          console.log(response);

          if (response.ok) {
            console.log(response.data);
            return response;
          } else {
            return response.json().then((data) => {
              let errorMessage = 'Authentication Failed';
              if (data && data.error && data.error.message) {
                errorMessage = data.error.message;
              }
              //   console.log(errorData);
              throw new Error(errorMessage);
            });
          }
        })
        .then((data) => {
          console.log('this is important', data);
          navigate('/headstepper');
        })
        .catch((error) => {
          console.log('There was an error', error);
        });
    }
  };

  const validateEmail = (email) => {
    return /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email);
  };

  const validatePhone = (contact_num) => {
    return /^[0-9]{10}$/.test(contact_num);
  };

  return (
    <div className="form1 sign-in">
      <form onSubmit={handleSubmit}>
        <h1>Sign in</h1>
        <div className="input-group1">
          <i className="bx fa fa-envelope "></i>
          <input
            type="text"
            placeholder="Email"
            className="input1"
            name="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />
          <br />
          {errors.email && (
            <span className="error" style={{ color: 'red' }}>
              {errors.email}
            </span>
          )}
        </div>
        <div className="input-group1">
          <i className="bx fa fa-user"></i>
          <input
            type="text"
            placeholder="Phone"
            className="input1"
            name="contact"
            value={contact_num}
            onChange={(e) => setContactNum(e.target.value)}
          />
          <br />
          {errors.contact_num && (
            <span className="error" style={{ color: 'red' }}>
              {errors.contact_num}
            </span>
          )}
        </div>
        <div className="input-group1">
          <i className="bx fa fa-lock"></i>
          <input
            type="password"
            placeholder="Password"
            className="input1"
            name="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
          <br />
          {errors.password && (
            <span className="error" style={{ color: 'red' }}>
              {errors.password}
            </span>
          )}
        </div>
        <button className="btn-1" type="submit">
          Sign in
        </button>
        <Link to="/Forgotpass" className="para">
          <b>Forgot password?</b>
        </Link>
        <p className="para">
          <span>Don't have an account ?</span> {'  '}
          <b onClick={toggle} className="pointer1">
            Sign up here
          </b>
        </p>
      </form>
    </div>
  );
};

export default Login;
