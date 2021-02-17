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
import CardGiftcardIcon from '@material-ui/icons/CardGiftcard';
import axios from 'axios';

const styles = theme => ({
  root: {
    maxWidth: 250,
  },
});

class UserGift extends  Component{

  state= {enable: true}
  constructor(props){
    super(props);
    this.interest = (this.props.interest - 1) / 2;
  }


    handleClickGift = () =>{
      console.log('clicked');
      axios.get('/user/' + this.props.username + '/list/' + this.props.listId + '/gift/' + this.props.giftId
      + '/buy').then(
        res => {
          console.log('success');
          window.location.reload(false);
        }
      ).catch(
        err => {
          console.log('failure');
        }
      )
    }

    
  render(){
    const {classes} = this.props;
    let button;
    if (this.props.eb){
      button = <Button style={{marginRight:0}} onClick={this.handleClickGift} disabled>
      <CardGiftcardIcon color="secondary"></CardGiftcardIcon>
      </Button>
    }else{
      button = <Button style={{marginRight:0}} onClick={this.handleClickGift}>
      <CardGiftcardIcon color="secondary"></CardGiftcardIcon>
      </Button>
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
                {this.props.price} تومان
              </Typography>
            </CardContent>
          </CardActionArea>
          <CardActions>
          {/* <Link href={'//' + this.props.link} target="_blank" rel="noreferrer">
            لینک
          </Link> */}
            <Button href={'//' + this.props.link} size="small" color="primary">
                لینک
            </Button>
            {/* <Button size="small" color="secondary" onClick={this.deleteGift}>
              حذف
            </Button> */}
            {button}
            <Rating name="read-only" value={this.interest} readOnly size="small" style={{paddingRight:0}}/>
          </CardActions>
        </Card>
      );
  }
  
}

export default withStyles(styles)(UserGift)