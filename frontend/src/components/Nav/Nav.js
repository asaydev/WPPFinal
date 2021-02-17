import { TextField, Button } from "@material-ui/core";
import React, { Component } from "react";
import { Link } from "react-router-dom";
import SearchIcon from '@material-ui/icons/Search'

export default class Nav extends Component{

  handleLogout = () => {
    localStorage.clear();
    window.location.reload(false);
  }

  handleSearch = () => {
    console.log(this.searchPhrase)
  }
  
    render(){
      let buttons;
        if (localStorage.getItem('token')){
          buttons = (<ul className="nav justify-content-end">
            <li className="nav-item">
              <TextField onChange={e => this.searchPhrase = e.target.value} placeholder="جستجو... ">
              </TextField>
              <Button onClick={this.handleSearch}>
                  <SearchIcon></SearchIcon>
                </Button>
            </li>
          <li className="nav-item">
            <Link className="nav-link" to={'/product'} >کالاها</Link>
          </li>
          <li className="nav-item">
            <Link className="nav-link" to={'/profile'} >پروفایل</Link>
          </li>
          <li className="nav-item">
            <Link className="nav-link" to={'/'} onClick={this.handleLogout}>خروج</Link>
          </li>
        </ul>)

        }else{
            buttons = (<ul className="nav justify-content-end">
            <li className="nav-item">
              <Link className="nav-link" to={'/register'}>ثبت نام</Link>
            </li>
            <li className="nav-item">
              <Link className="nav-link" to={'/login'}>ورود</Link>
            </li>
            
          </ul>)
        }
        return (<nav className="navbar nabvar-expand navbar-light fixed-top">
        <div className="container">
          <Link className="navbar-brand" to={'/'}>صفحه اصلی</Link>
            <div className="nav justify-content-end">
              {buttons}
          </div>
        </div>
      </nav>)
    }
}