import React, { Component } from "react";

export default class Home extends Component{
    
    

    render(){
        console.log('home render');
        if (localStorage.getItem('token')){

            return <h2>سلام {localStorage.getItem('username')}</h2>
        }
        return <h2>خوش آمدید!</h2>
    }
}