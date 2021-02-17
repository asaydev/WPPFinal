import { Component } from "react";
import axios from 'axios';
import UserListWD from "./UserListWD";


export default class User extends Component{

    
    state = {lists: []};
    constructor(props){
        super(props);
        
        axios.get('/user/list').then(
            res => {
                console.log(res.data.lists);
                // localStorage.setItem('lists', res.data)
                // console.log(typeof(res.data));
                this.setState({lists: res.data.lists});
            }
        ).catch(
            err => {
                console.log(err);
            }
        )

    }

    render(){
        // console.log(this.lists)
        return(<div style={{width: '100%', padding:80}}>
            {
                this.state.lists.map((list) => (
                    <UserListWD name={list.name} gifts={list.gifts.length} listId={list.id} username={this.props.match.params.username}></UserListWD>
                ))  
            }
        </div>)
    }
}