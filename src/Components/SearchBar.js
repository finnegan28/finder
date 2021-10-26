import React, { Component } from "react";
import axios from "axios";


class SearchBar extends Component {
  constructor(props) {
    super(props)
    
    this.state = {
      search: ''
      }                  
  }
  
  handleChange = event =>{
    this.setState({[event.target.name]: event.target.value});
  }
  
  handleSubmit = event => {
    event.preventDefault();

    axios
      .post("http://127.0.0.1:8080/find", this.state) 
      .then(response=>{
        console.log(response);
        this.setState({ data: response.data });
      })
      .catch((error) => {
        console.log(error);
      });
  };

  render() {
  const { search } = this.state  
    return (
      <div>
        <form onSubmit={this.handleSubmit}>
          <input
            placeholder="Search for music 'Artist - Track'"
            type="text" 
            name="search" 
            value={search}
            onChange={this.handleChange}/>
          <button type="submit">Search</button>
        </form>
                
        <div>
        {this.state.data}
        </div>
      </div>
    );
  }
}

export default SearchBar;
