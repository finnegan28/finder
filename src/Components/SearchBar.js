import React, { Component } from "react";
import axios from "axios";


class SearchBar extends Component {
  state = {
    search: ""
  };

  onSearchChange = e => {
    this.setState({
      search: e.target.value
    });
  };

  handleSubmit = (event) => {
    event.preventDefault();
    const data = {
      search: this.state.search
    };
    axios
      .post("http://127.0.0.1:8080/find", data) 
      .then((response) => {
        console.log(response.data);
        this.setState({ data: response.data });
      })
      .catch((error) => {
        console.log(error);
      });
  };

  render() {
    return (
      <div className="post">
        <form className="post" onSubmit={this.handleSubmit}>
          <input
            placeholder="Search for music 'Artist - Track'" value={this.state.search}
            onChange={this.onSearchChange} required
          />
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