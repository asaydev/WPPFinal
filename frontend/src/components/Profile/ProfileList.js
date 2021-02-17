import { Button, TextField } from '@material-ui/core';
import { withStyles } from '@material-ui/core/styles';
import React, { Component } from 'react';
import axios from 'axios';

const styles = theme => ({
    root: {
      '& > *': {
        margin: theme.spacing(1),
      },
    },
    textField: {
        paddingLeft: theme.spacing(3)
    },
    button: {
        marginLeft: theme.spacing(3)
    }
  });


class ProfileList extends Component {

    createNewList = e => {
        console.log('here');
        const data = {name: this.newListName};
        console.log(data)
        axios.post('/user/list', data).then(
            res => {
                console.log(res);
            }
        ).catch(
            err => {
                console.log(err);
            }
        )
    }

    render(){
        const {classes} = this.props;
    return (
        <div>
            <p>‌شما می‌توانید از این قسمت لیست‌هایتان را مشاهده کنید یا لیست جدید بسازید.</p>
            <TextField onChange={e => this.newListName = e.target.value} className={classes.textField} size="small"></TextField>
            <Button className={classes.button} variant="contained" color="primary" onClick={this.createNewList}>ایجاد لیست جدید</Button>
            <Button variant="contained" href="/list">مشاهده لیست‌ها</Button>
        </div>
    )
    }
}

export default withStyles(styles)(ProfileList)