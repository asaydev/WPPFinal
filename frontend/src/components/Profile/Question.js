import { Component } from "react";
import Product from "../Product/Product";

class Question extends Component{
    render() {
        return(
            <div>
            <Product name={this.props.name} image={this.props.image} link={this.props.link} brand={this.props.brand}
            category={this.props.category} price={this.props.price}></Product>
            </div>
        )
    }
}

export default Question;