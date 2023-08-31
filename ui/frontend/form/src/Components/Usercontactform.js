import React,{useState,useEffect,useContext} from 'react'
import {v4 as uuidv4} from 'uuid'
import Cookies from 'js-cookie'
import './Usercontactform.css';
import { MultiStepContext } from '../StepContext';

const Usercontactform = () => {
    const { setCurrentStep, userData, SetUserData } =
    useContext(MultiStepContext);

    const storedFullName = decodeURIComponent(Cookies.get('full_name') || '');

    const [full_name,setfull_name] = useState('');
    const [applicant_contact, setapplicant_contact] = useState('');
    const [loan_type,setloan_type] = useState('');
    const [errors, setErrors] = useState({});
    const [device_id, setdeviceID] = useState('');

    useEffect(() => {
        let storedID = Cookies.get('device_id');
        if (!storedID) {
            storedID = uuidv4();
            Cookies.set('device_id', storedID, { expires: 2 });
        }
        setdeviceID(storedID);

        let token = Cookies.get('token');
        if (!Cookies.get('token')) {
            Cookies.set('token', '', { expires: 2 });
        }
    }, []);

    const handleSubmit = async(e) => {
        e.preventDefault();
        let formErrors = {};

        if(!full_name || full_name.length < 3) {
            formErrors.full_name = "Name must be atleast 3 character long!"; 
        }

        const phoneRegex = /^[0-9]{10}$/;
        if(!phoneRegex.test(applicant_contact)) {
            formErrors.aplicant_contact = "Invalid phone number";
        }

        if(!loan_type) {
            formErrors.loan_type = "please select a loan type"
        }

        setErrors(formErrors);

     
        if (Object.keys(formErrors).length === 0) {
            Cookies.set('full_name', full_name, { expires: 2 });
            Cookies.set('applicant_contact', applicant_contact, { expires: 2 });
            Cookies.set('loan_type', loan_type, { expires: 2 });

            try {
                const response = await fetch('/proxy?url=http://127.0.0.1:8080/loan/insert', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({ device_id,full_name, applicant_contact, loan_type }),
                });
    
                const data = await response.json();
    
                if (response.ok) {
                    console.log('Form submitted successfully:', data);
                    setCurrentStep(2);
                } else {
                    console.error('API error:', data);
                }
            } catch (error) {
                console.error('There was an error submitting the form:', error);
            }
        } else {
            console.log("Form has errors.");
        }
    };
  return (
    <div className='length'>
    <div className="wrapper py-5">
        <h2>User Details</h2>
        <form onSubmit={handleSubmit}>
            <div className="form-group">
                <label htmlFor="name">Full Name</label>
                <input type="text" name="Name" id="name" placeholder="First and Last"
                    value={full_name}
                    onChange={(e) => setfull_name(e.target.value)}
                 minlength="3" maxlength="25" />
                {errors.full_name && <span className="error" style={{ color: 'red' }}>
              {errors.full_name}
            </span>}
            </div>
            <div className="form-group">
                <label htmlFor="contactNo">Contact No</label>
                <input type="tel" name="ContactNo" id="contactNo" placeholder="e.g., 1234567890"
                    value={applicant_contact}
                    onChange={(e) => setapplicant_contact(e.target.value)} /><br/>
                {errors.aplicant_contact && <span className="error" style={{ color: 'red' }}>
              {errors.aplicant_contact}
            </span>}
            </div>
            <div className="form-group">
                <label htmlFor="lang">Loan type</label>
                <select name="languages" id="lang" value={loan_type} onChange={(e) => setloan_type(e.target.value)}>
                    <option value="" disabled>Loan_Type</option>
                    <option value="Home">Home Loan</option>
                    <option value="Personal">Personal Loan</option>
                    <option value="Business">Business Loan</option>
                    <option value="credit">Credit Card</option>
                    <option value="wheeler">Two Wheeler Loan</option>
                    <option value="car">Car Loan</option>
                    <option value="Mortgage">Mortgage Loan</option>
                    <option value="Education">Education Loan</option>
                    <option value="Gold">Gold Loan</option>
                </select>
                {errors.loan_type &&  <span className="error" style={{ color: 'red' }}>
              {errors.loan_type}
            </span>}
            </div>
            <div className="form-group">
                <button type="submit" className="mybtn silver" onClick={handleSubmit}>Next<i className="fa fa-angle-double-right"></i></button>
            </div>
        </form>
    </div>
</div>


  )
}

export default Usercontactform
