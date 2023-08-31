import React from 'react';
import { Link } from 'react-router-dom';

function Mail() {
  return (
    <div
      style={{
        display: 'flex',
        justifyContent: 'center',
        marginTop: '70px',
      }}
    >
      <div
        class="card app font-sans min-w-screen min-h-screen bg-grey-lighter py-8 px-4"
        style={{ width: '30rem' }}
      >
        <img
          src="/assets/530.jpg"
          class="card-img-top rounded mx-auto d-block"
          alt="..."
          style={{ width: '200px' }}
        />
        <div class="card-body">
          <h1 class="card-title text-center">E-mail Confirmation</h1> <br />
          <hr />
          <br /> <br />
          <b>
            <p class="card-text">
              Hey, <br />
              <br />
              It looks like you just signed up for The App, that’s awesome! Can
              we ask you for email confirmation? Just click the button bellow.
            </p>
          </b>
          <br />
          <br />
          <Link to="/" class="btn btn-primary w-100">
            CONFIRM EMAIL ADDRESS
          </Link>
        </div>
        <br />
        <div class="content__footer mt-8 text-center text-grey-darker">
          <h3 class="text-base sm:text-lg mb-4">Thanks for using The App!</h3>
        </div>
      </div>
    </div>
  );
}

export default Mail;
