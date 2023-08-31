import React from 'react'
import './Navbar.css'

 const Navbar = () => {
  return (
    <div>
      <nav class="navbar navbar-expand-lg navbar-transparent bg-transparent">
        <div className='container'>
  {/* <a class="navbar-brand fs-2" style={{color:"black",fontFamily:"cursive"}}>Loan-Application</a> */}
  <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
    <span class="navbar-toggler-icon"></span>
  </button>

  <div class="collapse navbar-collapse" id="navbarSupportedContent">
    <ul class="navbar-nav ms-auto">
      
    </ul>
    <form class="form-inline my-4  my-lg-0" style={{marginRight:"40px",display:"flex"}}>
      {/* <input class="form-control mr-sm-2" type="search" placeholder="Search" aria-label="Search"/> */}
      <button className="mybtn silver me-4 col-lg-4 col-md-4 col-sm-12 col-xs-12" type="submit" style={{margin:"5px"}}>Sign-in</button>
            <button className="mybtn silver  me-4 col-lg-4 col-sm-12 col-xs-12" type="submit" style={{margin:"5px"}}>Sign-up</button>
            <button className="mybtn silver  me-4 col-lg-4 col-sm-12 col-xs-12" type="submit" style={{margin:"5px"}}>Logout</button>
    </form>
    </div>
  </div>
</nav>
    </div>
  )
}
export default Navbar
