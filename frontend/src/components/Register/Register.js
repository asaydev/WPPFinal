import React, { Component } from "react";
import axios from 'axios';

export default class Register extends Component{

    handleSubmit = e => {
        e.preventDefault();
        const data = {
            firstname: this.firstName,
            username: this.username,
            password:this.password

        }
        console.log(data);
        axios.post('/user/register', data).then(
            res => {
                console.log(res.data)
                localStorage.setItem('username', this.firstName)
            }
        ).catch(
            err => {
                console.log(err)
            }
        )
    }
    render(){
        return(
            <div className="auth-wrapper">
                <div className="auth-inner">
                    <form onSubmit={this.handleSubmit}>
                        <h3> ثبت نام </h3>

                        <div className="form-group">
                            {/* <label>Name</label> */}
                            <input type="text" className="form-control" placeholder="نام" 
                            onChange={e => this.firstName = e.target.value}/>
                        </div>
                        <div className="form-group">
                            {/* <label>Username</label> */}
                            <input type="text" className="form-control" placeholder="نام کاربری"
                            onChange={e => this.username = e.target.value}/>
                        </div>
                        <div className="form-group">
                            {/* <label>Password</label> */}
                            <input type="password" className="form-control" placeholder="گذرواژه"
                            onChange={e => this.password = e.target.value}/>
                        </div>
                        <button className="btn btn-primary btn-block">ثبت نام</button>
                    </form>
                </div>
            </div>
        )
    }
}