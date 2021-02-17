import React, { Component } from 'react';
import { withStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActionArea from '@material-ui/core/CardActionArea';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import CardMedia from '@material-ui/core/CardMedia';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';
import Rating from '@material-ui/lab/Rating';
import axios from 'axios';

const styles = theme => ({
  root: {
    maxWidth: 250,
  },
});

class Gift extends  Component{

    constructor(props){
      super(props);
      this.interest = (this.props.interest) / 2;
    }
    

    deleteGift = () => {
        axios.delete('/user/list/' + this.props.listId + '/gift/' + this.props.giftId).then(
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
    let text;
    if (this.props.eb){
      text = 'خریده شده!';
    }else{
      text = this.props.price  + ' تومان '
    }
    return (
        <Card className={classes.root}>
          <CardActionArea>
            <CardMedia
              component="img"
              alt="Contemplative Reptile"
              height="140"
              image="https://material-ui.com/static/images/cards/contemplative-reptile.jpg"
              title="Contemplative Reptile"
            />
            <CardContent>
              <Typography gutterBottom variant="h5" component="h2">
                {this.props.name}
              </Typography>
              <Typography variant="body2" color="textSecondary" component="p">
                {text}
              </Typography>
              {/* <Typography vatiant="body2" color="textSecondary" component="p">
                {text}
              </Typography> */}
            </CardContent>
          </CardActionArea>
          <CardActions>
          {/* <Link href={'//' + this.props.link} target="_blank" rel="noreferrer">
            لینک
          </Link> */}
            <Button href={'//' + this.props.link} size="small" color="primary">
                لینک
            </Button>
            <Button size="small" color="secondary" onClick={this.deleteGift}>
              حذف
            </Button>
            {/* <Button size="small" color="primary">
              {this.props.interest}/10
            </Button> */}
            {/* <Box component="fieldset" mb={3} borderColor="transparent"> */}
              <Rating name="read-only" value={this.interest} readOnly size="small"/>
            {/* </Box> */}
          </CardActions>
        </Card>
      );
  }
  
}

export default withStyles(styles)(Gift)