import axios from "axios";
import { Component } from "react";
import Product from './Product';
import { Grid, Typography} from '@material-ui/core';
import Pagination from '@material-ui/lab/Pagination';
import TextField from '@material-ui/core/TextField';
import SearchIcon from '@material-ui/icons/Search';
import PrimarySearchAppBar from './PrimarySearchAppBar';



class Products extends Component{

    state = {products: [], pages: 10}

    constructor(props){
        super(props);
        
        const query = new URLSearchParams(this.props.location.search);
        this.page = parseInt(query.get('page') || '1');
        // this.setState({page:parseInt(query.get('page') || '1')})
        console.log('page');
        // console.log.apply(this.state.page)
        // console.log(page)
        axios.get('/product?page=' + this.page).then(
            res => {
                console.log(res.data);
                this.setState({pages: res.data.meta.pages})
                this.setState({products: res.data.list})
            }
        ).catch(
            err => {
                console.log(err);
            }
        )
    }

    handleChange = (event, value) => {
        console.log(value);
        // this.setState({page: value})
        this.props.history.push('/product?page=' + value);
        window.location.reload(false);
    }

    render(){
        return(<main style={{padding: 80}}>
            {/* <PrimarySearchAppBar></PrimarySearchAppBar> */}
            <Grid container justify="center" spacing={4}>
        {
            this.state.products.map((product) => (
                <Product name={product.name} brand={product.brand} category={product.category} image={product.image}
                link={product.link} price={product.price}></Product>
            ))
        }
        </Grid>
        <Pagination style={{margin:80, marginTop:20}} count={this.state.pages} page={this.page} onChange={this.handleChange} color="secondary" />
    </main>
        )
    }
}

export default Products;