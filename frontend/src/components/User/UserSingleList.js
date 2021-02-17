import { Grid, Typography} from '@material-ui/core';
import axios from 'axios';
import React, { Component } from 'react'
import UserGift from './UserGift'

export default class UserSingleList extends Component {
    state = {list: {gifts:[]}}

    constructor(props){
        super(props);
        axios.get('/user/'+ this.props.match.params.username + '/list/' + this.props.match.params.id).then(
            res => {
                this.setState({list: res.data});
                console.log(this.state.list.gifts)
            }
        ).catch(
            err => {
                console.log(err);
            }
        )
    }

    render(){
        return(<main style={{padding:80}}>
            <Typography style={{paddingBottom: 30}}>لیست {this.state.list.name}</Typography>
            <Grid container justify="center" spacing={4}>
        {   
            this.state.list.gifts.map((gift) => (
                <UserGift eb={gift.eb} username={this.props.match.params.username} listId={this.props.match.params.id} name={gift.name} price={gift.price} interest={gift.interest} giftId={gift.id} link={gift.link}></UserGift>
                // console.log(gift)
            ))
        }
        </Grid>
    </main>)
    }
}

// export default SingleList
