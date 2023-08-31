import React, { useState, createContext } from 'react';
import App from './App';

export const MultiStepContext = createContext();

const StepContext = () => {
  const [currentStep, setCurrentStep] = useState(1);
  const [userData, SetUserData] = useState({});
  const [finalData, SetFinalData] = useState({});

  return (
    <div>
      <MultiStepContext.Provider value={{ currentStep, setCurrentStep, userData, SetUserData, finalData, SetFinalData}}>
        <App/>
      </MultiStepContext.Provider>
    </div>
  );
}

export default StepContext;















// import React, { useState, createContext, useCallback } from 'react';
// import App from './App';

// export const saveData = async (userData) => {
//   console.log('Sending user data:', userData);
//   try {
  
//     const response = await fetch('proxy?url=http://127.0.0.1:8080/loan/insert', {
//       method: 'POST',
//       headers: {
//         'Content-Type': 'application/json',
//       },
      
//       body: JSON.stringify({
//         loan_type: userData.loan_type,
//         loan_amount: parseFloat(userData.loanAmount),
//         pincode: parseInt(userData.pincode),
//         employment_type: userData.employment_type,
//         gross_monthly_income: parseFloat(userData.GrossMonthly),
//         tenure : parseInt(userData.tenure)
//         // add any other necessary fields
//       }),
//     });
//     console.log("Response status:", response.status);

//     if (!response.ok) {
//       const message = await response.text();
//       console.error("Response body:", message);
//       throw new Error(message);
//     }

//   } catch (error) {
//     console.error("Error in saveData:", error);
//     throw error;
//   }
// };


// export const MultiStepContext = createContext();

// const StepContext = () => {
//   const [currentStep, setCurrentStep] = useState(1);
//   const [userData, SetUserData] = useState({});
//   const [finalData, SetFinalData] = useState({});

//   const saveDataCaLLBack = useCallback(saveData, []);

//   return (
//     <div>
//       <MultiStepContext.Provider value={{ currentStep, setCurrentStep, userData, SetUserData, finalData, SetFinalData, saveData: saveDataCaLLBack }}>
//         <App/>
//       </MultiStepContext.Provider>
//     </div>
//   );
// }

// export default StepContext;
