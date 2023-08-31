import React, { useContext} from 'react'
import { MultiStepContext } from '../StepContext';

const StepSeven = () => {
    const {setCurrentStep , userData, SetUserData} = useContext(MultiStepContext);

    function getCookie(name) {
      let value = "; " + document.cookie;
      let parts = value.split("; " + name + "=");
      if (parts.length === 2) return parts.pop().split(";").shift();
  }

 
  const handleNext = async () => {
    if (userData['tenure'] && userData['tenure'].trim() !== '') {
        const userId = getCookie('loan_id');
        if (!userId) {
            alert('User ID not found in cookies.');
            return;
        }

        const apiUrl = `proxy?url=http://127.0.0.1:8080/loan/update/${userId}`;
        try {
            const response = await fetch(apiUrl, {
                method: 'PATCH',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    tenure: parseInt(userData['tenure'])
                })
            });

            if (!response.ok) {
                throw new Error('Network response not ok');
            }

            setCurrentStep(6);
        } catch (error) {
            console.error('There has been a problem with your fetch operation:', error);
            alert("An error occurred while saving data. Please try again.");
        }
    } else {
        alert("Please enter a valid tenure.");
    }
};
  return (
    <div>
        <section>
        <div className="container d-flex justify-content-center align-items-center mt-5 box">
          <div className="row">
            <div className="card d-flex shadow-lg " style={{ backgroundColor: '#F7F8FA' }}>
              <div className="card-body">
                <div className="container text-center my-3">
                  <h2>
                    <b>Loan Application</b>
                  </h2>
                </div>

                <div className="firstbox w-100">
            <h2 className="fs-4">Step 6 : tenure</h2>
            <div className="row">
            </div>

            <div className="form-group col-md-12 mt-3 hoverEffect">
              <label> Enter tenure:</label>
              <input
                type="text"
                name="test"
                id="cb15"
                placeholder="Enter tenure"
                className="form-control"
                value={userData['tenure']}
                onChange={(e) => SetUserData({...userData, 'tenure': e.target.value})}
              />
            </div>

            <div className="d-flex">
              <button
                className="btn btn-success m-3 mt-5 w-50"
                onClick={()=>setCurrentStep(6)}
              >
                Previous
              </button>
            
                <button className="btn btn-success m-3 mt-5 w-50" onClick={handleNext} > Next </button>
            </div>
          </div>
          </div>
          </div>
          </div>
          </div>
          </section>
</div>
      
  )
}
export default StepSeven
