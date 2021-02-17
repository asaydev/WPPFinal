import { Component } from "react";
import Switch from '@material-ui/core/Switch';
import axios from "axios";

class EditUser extends Component{

    state = {checked: false}

    handleChange = () => {
        this.setState({checked: !this.state.checked})
    }

    handleSubmit = () => {
        const data = {
            firstName: this.name,
            lastName: this.lastName,
            password: this.password
        }
        axios.put('/user', data).then(
            res => {
                console.log('success-edit');
            }
        ).catch(
            err => {
                console.log(err);
            }
        )
    }


    render(){
        let form;
        if (this.state.checked){
            form = <form onSubmit={this.handleSubmit}>
            <div className="form-group">
                <input type="text" className="form-control" placeholder="نام"
                onChange={e => this.name = e.target.value}/>
            </div>
            <div className="form-group">
                <input type="text" className="form-control" placeholder="نام خانوادگی"
                onChange={e => this.lastName = e.target.value}/>
            </div>
            <div className="form-group">
                <input type="password" className="form-control" placeholder="گذرواژه"
                onChange={e => this.password = e.target.value}/>
            </div>
            <button className="btn btn-primary btn-block">ذخیره</button>
        </form>
        }
        return(<main>
            <body2>ویرایش اطلاعات</body2>
            <Switch color="primary" onChange={this.handleChange} checked={this.state.checked} name="edit" inputProps={{ 'aria-label': 'secondary checkbox' }}></Switch>
            {form}
        </main>)
    }
}

export default EditUser;