import { Component } from "react";
import axios from 'axios';
import UserList  from './UserList'


export default class UserLists extends Component{

    
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
                    <UserList name={list.name} gifts={list.gifts.length} listId={list.id}></UserList>
                ))
            }
        </div>)
    }
}