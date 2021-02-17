import { Button, withStyles } from "@material-ui/core";
import { Component } from "react";

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
        marginLeft: theme.spacing(12)
    }
  });

class Suggestion extends Component{
    
    render(){
        const {classes} = this.props;
        return (
            <div>
                <Button className={classes.button} variant="contained" color="primary">پیشنهادها</Button>
                <Button variant="contained" color="primary" href="/question" >سوالات</Button>
            </div>
        )
    }
}

export default withStyles(styles)(Suggestion);