import React from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import Combine from './Components/Combine';
import '../node_modules/bootstrap/dist/css/bootstrap.min.css';
import ForgotPass from './Components/forgetpass';
import HeadStepper from './Components/HeadStepper';
import Otp from './Components/Otp';
import NewPassword from './Components/NewPassword';
import Usercontactform from './Components/Usercontactform';
import Mail from './Components/Mail';
import { ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
// import Navbar from './Components/Navbar';

function App() {
  return (
    <div className="App">
      <BrowserRouter>
        <Routes>
        <Route path='/' element={<div className='App'><><HeadStepper/></></div>}/> 
          <Route path="/Combine" element={<Combine />} />
          <Route path="/Forgotpass" element={<ForgotPass />} />
          <Route path="/otp" element={<Otp />} />
          <Route path="/NewPassword" element={<NewPassword />} />
          <Route path="/mail" element={<Mail />} />
          {/* <Route path="/HeadStepper" element={<HeadStepper />} /> */}
        </Routes>
      </BrowserRouter>
      <ToastContainer />
    </div>
  );
}
export default App;
