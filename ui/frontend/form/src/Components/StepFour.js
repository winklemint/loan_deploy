import React, {useState , useContext} from 'react'
import { MultiStepContext } from '../StepContext';

 const StepFour = () => {
  const {setCurrentStep , userData, SetUserData} = useContext(MultiStepContext);


  const [value, setValue] = useState(userData.gross_monthly_income || ''); // Initial value of the slider
  const [Range, setRange] = useState('');

  const convertToWords = (num) => {
    const units = [
      '',
      'one',
      'two',
      'three',
      'four',
      'five',
      'six',
      'seven',
      'eight',
      'nine',
    ];
    const tens = [
      '',
      '',
      'twenty',
      'thirty',
      'forty',
      'fifty',
      'sixty',
      'seventy',
      'eighty',
      'ninety',
    ];
    const thousands = ['', 'thousand', 'lakh'];

    if (num === 0) return 'zero';

    const convertChunkToWords = (chunk) => {
      let chunkWords = '';

      const hundreds = Math.floor(chunk / 100);
      const remainder = chunk % 100;

      if (hundreds > 0) {
        chunkWords += units[hundreds] + ' hundred';
      }

      if (remainder > 0) {
        if (chunkWords !== '') {
          chunkWords += ' ';
        }

        if (remainder < 20) {
          chunkWords += units[remainder];
        } else {
          const tensPlace = Math.floor(remainder / 10);
          const onesPlace = remainder % 10;

          chunkWords += tens[tensPlace];

          if (onesPlace > 0) {
            chunkWords += '-' + units[onesPlace];
          }
        }
      }

      return chunkWords;
    };

    let words = '';
    let chunkIndex = 0;

    while (num > 0) {
      const chunk = num % 1000;
      if (chunk !== 0) {
        const chunkWords = convertChunkToWords(chunk);
        if (chunkIndex > 0) {
          words = chunkWords + ' ' + thousands[chunkIndex] + ' ' + words;
        } else {
          words = chunkWords;
        }
      }
      num = Math.floor(num / 1000);
      chunkIndex++;
    }
    return words;
  };

  function getCookie(name) {
    let value = "; " + document.cookie;
    let parts = value.split("; " + name + "=");
    if (parts.length === 2) return parts.pop().split(";").shift();
 }

  const handleChange = (event) => {
    setValue(event.target.value);
    SetUserData((prevUserData) => ({ ...prevUserData, gross_monthly_income: event.target.value }));
  };

  const handleNext = async () => {
    if (value) {
        // Extracting user ID from cookies
        const userId = getCookie('loan_id');
        if (!userId) {
            alert('User ID not found in cookies.');
            return;
        }

        // Define API URL
        const apiUrl = `proxy?url=http://127.0.0.1:8080/loan/update/${userId}`;

        try {
            // Make an API call to save the Gross Monthly Income
            const response = await fetch(apiUrl, {
                method: 'PATCH',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    gross_monthly_income: parseFloat(value)
                })
            });

            if (!response.ok) throw new Error('Network Response not ok');

            // Update local state and proceed to next step
            SetUserData(prevUserData => ({ ...prevUserData, gross_monthly_income: value }));
            setCurrentStep(6);
        } catch (error) {
            console.error('There has been a problem with your fetch operation:', error);
            alert("An error occurred while saving data. Please try again.");
        }
    } else {
        alert("Please select a gross monthly income.");
    }
};


  const words = convertToWords(value);
  return (
    <div>
        <section>
    <div className="container d-flex justify-content-center align-items-center mt-5 box">
      <div className="row">
        <div
          className="card d-flex shadow-lg "
          style={{ backgroundColor: '#F7F8FA' }}
        >
          <div className="card-body">
            <div className="container my-3">
              <h2>
                <b>Loan Application</b>
              </h2>
            </div>
            <div className="firstbox w-100">
            <h2 className="fs-4">Step 4: Gross Monthly Income</h2>
            <div className="row">
              <input
                // type="range"
                min={0}
                max={1000000}
                value={value}
                onChange={handleChange}
              />
              <p>Value: {value}</p>
              <p>Words: {words}</p>
            </div>

            <div className="d-flex">
              <button
                className="btn btn-success m-3 mt-5 w-50"
                onClick={() => setCurrentStep(3)}
              >
                Previous
              </button>
            
                <button
                  className="btn btn-success m-3 mt-5 w-50"
                  onClick={handleNext}
                >
                  Next
                </button>
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
export default StepFour
