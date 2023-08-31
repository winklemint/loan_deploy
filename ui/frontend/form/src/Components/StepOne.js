import React, { useState, useContext, useEffect } from 'react';
import './StepForm.css';
import { MultiStepContext } from '../StepContext';

const StepOne = () => {
  const { setCurrentStep, userData, SetUserData } =
    useContext(MultiStepContext);

  const [selectedOption, setSelectedOption] = useState(
    userData.loan_type || ''
  );
  const [isOptionSelected, setIsOptionSelected] = useState(false);

  const handleRadioChange = (event) => {
    setSelectedOption(event.target.value);
    setIsOptionSelected(true);
    const updatedData = { ...userData, loan_type: event.target.value };
    SetUserData(updatedData);
  };

  useEffect(() => {
    SetUserData({ ...userData, loan_type: selectedOption });
  }, [selectedOption]);

  function getCookie(name) {
    let value = "; " + document.cookie;
    let parts = value.split("; " + name + "=");
    if (parts.length === 2) return parts.pop().split(";").shift();
  }

  const handleNext = async () => {
    const userId = getCookie('id');
    if (!userId) {
      alert('User ID not found in cookies.');
      return;
    }
    const apiUrl = `proxy?url=http://127.0.0.1:8080/loan/update/${userId}`;
    
    try {
      const response = await fetch(apiUrl, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ loan_type: selectedOption }),
      });

      if (!response.ok) {
        throw new Error('Network response was not ok');
      }
      
      const data = await response.json();
      console.log('Data successfully sent:', data);
      SetUserData({ ...userData, loan_type: selectedOption });
      setCurrentStep(3);

    } catch (error) {
      console.error('There was a problem with the fetch operation:', error);
      alert('An error occurred while sending data. Please try again.');
    }
  };

  return (
    <div>
      <section >
        {/* <div className="new img-fluid"> */}
        <div className="container d-flex justify-content-center align-items-center mt-5 box">
          <div className="row">
            <div
              className="card d-flex "
              style={{ backgroundColor: 'white',border:'none' }}
            >
              <div className="card-body">
                <div className="container text-center my-3">
                  <h2>
                    <b>Loan Application</b>
                  </h2>
                </div>
                <div>
                  <div className="firstbox">
                    <h2 className="fs-4 ">Step 1: Purpose of Loan</h2>
                    <div className="row ">
                      <div className="form-group col-md-4 hoverEffect " style={{border:"none"}}>
                        <input
                          type="radio"
                          name="test"
                          id="cb1"
                          className="form-control"
                          value="BuyHome"
                          checked={selectedOption === 'BuyHome'}
                          onChange={handleRadioChange}
                        />
                        <label for="cb1" className="checked-label">
                          <img src="assets/home.png" alt="one" />
                          Buy Home:
                        </label>
                      </div>

                      <div className="form-group col-md-4  hoverEffect">
                        <input
                          type="radio"
                          name="test"
                          id="cb2"
                          className="form-control "
                          value="loanagainst"
                          checked={selectedOption === 'loanagainst'}
                          onChange={handleRadioChange}
                        />
                        <label for="cb2" className="checked-label">
                          <img src="assets/loan.png" alt="two" />
                          Loan against
                        </label>
                      </div>

                      <div className="form-group col-md-4 hoverEffect ">
                        <input
                          type="radio"
                          name="test"
                          id="cb3"
                          className="form-control "
                          value="balancetransfer"
                          checked={selectedOption === 'balancetransfer'}
                          onChange={handleRadioChange}
                        />
                        <label for="cb3" className="checked-label">
                          <img src="assets/home-address.png" alt="three" />
                          Balance Transfer:
                        </label>
                      </div>

                      <div className="form-group col-md-4 hoverEffect">
                        <input
                          type="radio"
                          name="test"
                          id="cb4"
                          className="form-control "
                          value="homeimprovement"
                          checked={selectedOption === 'homeimprovement'}
                          onChange={handleRadioChange}
                        />
                        <label for="cb4" className="checked-label">
                          <img src="assets/home-improve.png" alt="four" />
                          Home Improvement:
                        </label>
                      </div>

                      <div className="form-group col-md-4 hoverEffect ">
                        <input
                          type="radio"
                          name="test"
                          id="cb5"
                          className="form-control "
                          value="buyplot"
                          checked={selectedOption === 'buyplot'}
                          onChange={handleRadioChange}
                        />
                        <label for="cb5" className="checked-label">
                          <img src="assets/house.png" alt="fifth" />
                          Buy Plot and Construct:
                        </label>
                      </div>

                      <div className="form-group col-md-4 hoverEffect">
                        <input
                          type="radio"
                          name="test"
                          id="cb6"
                          className="form-control "
                          value="construct"
                          checked={selectedOption === 'construct'}
                          onChange={handleRadioChange}
                        />
                        <label for="cb6" className="checked-label">
                          <img src="assets/insurance.png" alt="sixth" />
                          Construct on Plot:
                        </label>
                      </div>

                      <div className="form-group col-md-4 hoverEffect ">
                        <input
                          type="radio"
                          name="test"
                          id="cb7"
                          className="form-control "
                          value="commercial"
                          checked={selectedOption === 'commercial'}
                          onChange={handleRadioChange}
                        />
                        <label for="cb7" className="checked-label">
                          <img src="assets/leader.png" alt="seventh" />
                          Buy Commercial Property:
                        </label>
                      </div>
                    </div>

                    <div className="text-center" style={{ color: 'grey' }}>
                      <p>Select purpose of Loan</p>
                    </div>

                    <div className="text-center">
                      <button className="btn btn-success" onClick={handleNext}>
                        Next
                      </button>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        {/* </div> */}
      </section>
    </div>
  );
};
export default StepOne;
