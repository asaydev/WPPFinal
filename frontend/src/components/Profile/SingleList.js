import { Grid, Typography} from '@material-ui/core';
import axios from 'axios';
import React, { Component } from 'react'
import Gift from './Gift'

export default class SingleList extends Component {
    state = {list: {gifts:[]}}

    constructor(props){
        super(props);
        axios.get('/user/list/' + this.props.match.params.id).then(
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
                <Gift eb={gift.eb} listId={this.props.match.params.id} name={gift.name} price={gift.price} interest={gift.interest} giftId={gift.id} link={gift.link}></Gift>
                // console.log(gift)
            ))
        }
        </Grid>
    </main>)
    }
}

// export default SingleList
