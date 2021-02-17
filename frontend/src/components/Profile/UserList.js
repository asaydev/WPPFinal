import React, { Component } from 'react';
import { withStyles } from '@material-ui/core/styles';
import { Button } from '@material-ui/core';
import Accordion from '@material-ui/core/Accordion';
import AccordionSummary from '@material-ui/core/AccordionSummary';
import AccordionDetails from '@material-ui/core/AccordionDetails';
import Typography from '@material-ui/core/Typography';
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';
import axios from 'axios';

const styles = theme => ({
    root: {
      '& > *': {
        margin: theme.spacing(1),
      },
    },
    textField: {
        paddingLeft: theme.spacing(40)
    },
    button: {
        marginLeft: theme.spacing(1),
    }
  });


class UserList extends Component{

    constructor(props){
        super(props);
        this.name = props.name;
    }


    deleteList = () => {
        axios.delete('/user/list/' + this.props.listId).then(
            res => {
                console.log('deleted');
                window.location.reload(false);
            }
        ).catch(
            err => {
                console.log(err);
            }
        )
    }


    render(){
        const {classes} = this.props;
        return(
            <div >
                <Accordion>
                    <AccordionSummary
                    expandIcon={<ExpandMoreIcon />}
                    aria-controls="panel1a-content"
                    id="panel1a-header"
                    >
                        <Typography >{this.props.name}</Typography>
                    </AccordionSummary>
                    <AccordionDetails>
                        <Typography className={classes.textField}>تعداد آیتم‌ها: {this.props.gifts}</Typography>
                        <Typography className={classes.textField}>تعداد آیتم‌های در انتظار خرید: 2</Typography>
                        <Button className={classes.button} variant="contained" color="primary" href={"/list/" + this.props.listId}>مشاهده</Button>
                        <Button className={classes.button} variant="contained" color="secondary" onClick={this.deleteList}>حذف</Button>
                    </AccordionDetails>
                </Accordion>
            </div>
        )
    }

}

export default withStyles(styles)(UserList)