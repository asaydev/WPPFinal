import { Grid } from "@material-ui/core";
import axios from "axios";
import { Component } from "react";
import Question from './Question';

class Questions extends Component{
    state = {prods: []}

    constructor(props){
        super(props);
        axios.get('/user/question').then(
            res => {
                this.setState({prods: res.data})
                console.log(this.state.prods)
            }
        ).catch(
            err => {
                console.log(err);
            }
        )
    }

    render(){
        return (
            <main style={{padding: 80}}>
                <Grid container justify="center" spacing={4}>
                {
                    this.state.prods.map((prod) => (
                        <Question name={prod.name} price={prod.price} brand={prod.brand} category={prod.category}
                         image={prod.image} link={prod.link}></Question>
        ))
                }
                </Grid>
            </main>
        )
    }
}

export default Questions;