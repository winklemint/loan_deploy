import React, { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';

const Signup = ({ toggle }) => {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [user_contact_num, setContact] = useState('');
  const [password, setPassword] = useState('');
  const [errors, setErrors] = useState({});
  const navigate = useNavigate();

  const handleSubmit = (e) => {
    e.preventDefault();

    let formErrors = {};
    if (!name) {
      formErrors.name = 'Please enter your name';
    }
    if (!email) {
      formErrors.email = 'Please enter your email';
    } else if (!validateEmail(email)) {
      formErrors.email = 'Please enter a valid email';
    }
    if (!password) {
      formErrors.password = 'Please enter your password';
    } else if (password.length < 8) {
      formErrors.password = 'Password must contain at least 8 characters';
    }
    if (!user_contact_num) {
      formErrors.user_contact_num = 'Please enter your phone contact';
    } else if (!validatePhone(user_contact_num)) {
      formErrors.user_contact_num = 'Please enter a valid phone number';
    }

    setErrors(formErrors);

    if (
      Object.keys(formErrors).length === 0 &&
      name &&
      email &&
      user_contact_num &&
      password
    ) {
      fetch('/proxy?url=http://127.0.0.1:8080/user/add', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          name: name,
          email: email,
          user_contact_num: parseInt(user_contact_num),
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
              throw new Error(errorMessage);
            });
          }
        })
        .then((data) => {
          console.log('Registration Successfull', data);
          navigate('/mail');
        })
        .catch((error) => {
          console.error('Error during registration', error);
        });
    }
  };

  const validateEmail = (email) => {
    return /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email);
  };

  const validatePhone = (user_contact_num) => {
    const indianMobileRegex = /^[6789]\d{9}$/;
    return indianMobileRegex.test(user_contact_num);
  };

  return (
    <div className="form1 sign-up">
      <form onSubmit={handleSubmit}>
        <h1>Create Account</h1>
        <div className="input-group1">
          <i className="bx fa fa-user"></i>
          <input
            type="text"
            placeholder="Name"
            className="input1"
            value={name}
            name="name"
            onChange={(e) => setName(e.target.value)}
          />
          <br />
          {errors.name && (
            <span className="error" style={{ color: 'red' }}>
              {errors.name}
            </span>
          )}
        </div>
        <div className="input-group1">
          <i className="bx fa fa-envelope"></i>
          <input
            type="email"
            placeholder="Email"
            className="input1"
            value={email}
            name="email"
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
          <i className="bx fa fa-phone"></i>
          <input
            type="tel"
            placeholder="Phone"
            className="input1"
            value={user_contact_num}
            name="contact"
            onChange={(e) => setContact(e.target.value)}
          />
          <br />
          {errors.user_contact_num && (
            <span className="error" style={{ color: 'red' }}>
              {errors.user_contact_num}
            </span>
          )}
        </div>
        <div className="input-group1">
          <i className="bx fa fa-lock"></i>
          <input
            type="password"
            placeholder="Password"
            className="input1"
            value={password}
            name="password"
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
          Sign up
        </button>
        <p className="para">
          <span>Already have an account?</span>
          {'  '}
          <b onClick={toggle} className="pointer1">
            Sign in here
          </b>
        </p>
      </form>
    </div>
  );
};

export default Signup;
