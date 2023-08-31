import React, { useState, useEffect } from 'react';
import Login from './login';
import Signup from './signup';
import '../Components/Combine.css';

function Combine() {
  const [isSignIn, setIsSignIn] = useState(false);

  const toggle = () => {
    setIsSignIn(!isSignIn);
  };

  // Automatically switch to the "Sign in" view after a delay
  useEffect(() => {
    setTimeout(() => {
      setIsSignIn(true);
    }, 200);
  }, []);

  return (
    <div>
      <div
        id="container1"
        className={`container1 ${isSignIn ? 'sign-in' : 'sign-up'}`}
      >
        <div className="row1 content-row1">
          <div className="col1 align-items-center1 flex-col">
            <div className="text1 sign-in">
              {/* <h1 className="head2">Welcome</h1> */}
              <h1 className="head2" style={{ fontSize: '50px' }}>
                Hello, Friend!
              </h1>
              <p className="para" style={{ fontSize: '32px' }}>
                Enter your personal details and start your journey with us
              </p>
            </div>
            <div className="img1 sign-in"></div>
          </div>

          <div className="col1 align-items-center1 flex-col">
            <div className="img1 sign-up"></div>
            <div className="text1 sign-up">
              <h1 className="head2" style={{ fontSize: '60px' }}>
                Welcome
              </h1>
              <h2 className="head2" style={{ fontSize: '32px' }}>
                Join with us
              </h2>
              <p className="para" style={{ fontSize: '27px' }}>
                To keep connected with us please login with your personal info
              </p>
            </div>
          </div>
        </div>
        <div className="row1">
          <div className="col1 align-items-center1 flex-col sign-up">
            <Signup toggle={toggle} />
          </div>

          <div className="col1 align-items-center1 flex-col sign-in">
            <Login toggle={toggle} />
          </div>
        </div>
      </div>
    </div>
  );
}

export default Combine;
