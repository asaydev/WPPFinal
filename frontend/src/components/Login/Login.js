import { Component } from "react";
import axios from 'axios';

export default class Login extends Component{
    handleSubmit = e => {
        e.preventDefault();
        const data = {
            username: this.username,
            code: "1357"
        }
        console.log(data);
        axios.post('/login', data).then(
            res => {
                console.log(res);
                localStorage.setItem('token', res.data.access_token);
                localStorage.setItem('username', this.username);

                this.props.history.push('/profile');
                window.location.reload(false);


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
                        <h3>ورود</h3>
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
                        <button className="btn btn-primary btn-block">ورود</button>
                    </form>
                </div>
            </div>
        )
    }
}