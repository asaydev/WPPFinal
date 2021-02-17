import React, { Component } from 'react';
import { withStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActionArea from '@material-ui/core/CardActionArea';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import CardMedia from '@material-ui/core/CardMedia';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';
import DialogSelect from './DialogSelect';

const styles = theme => ({
  root: {
    maxWidth: 250,
  },
});

class Product extends  Component{
  

    
  render(){
    const {classes} = this.props;
    return (
        <Card className={classes.root}>
          {/* <CardActionArea> */}
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
              <Typography variant="caption" color="textSecondary" component="p">
                دسته {this.props.category} 
              </Typography>
              <Typography variant="caption" color="textSecondary" component="p">
                 برند {this.props.brand} 
              </Typography>
            </CardContent>
          {/* </CardActionArea> */}
          <CardActions>
            <Button href={'//' + this.props.link} size="small" color="primary" >
                لینک
            </Button>
            {/* <Button style={{marginRight: 100}}>
              <AddIcon color="primary"></AddIcon>
            </Button> */}
            <Button style={{marginRight: 100}}>
              <DialogSelect name={this.props.name} brand={this.props.brand} link={this.props.link}
              price={this.props.price}></DialogSelect>
            </Button>
            
            {/* <Button size="small" color="primary">
              {this.props.interest}/10
            </Button> */}
          </CardActions>
        </Card>
      );
  }
  
}

export default withStyles(styles)(Product)