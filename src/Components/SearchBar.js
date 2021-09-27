import React from 'react';

const SearchBar = ({keyword,setKeyword}) => {
  const BarStyling = {width:"20rem",background:"#F2F1F9", border:"none", padding:"0.5rem"};
  const BtnStyling = {width:"10rem",background:"#F2F1F9", border:"none", padding:"0.5rem"};
  return (
  <form>
    <input 
     style={BarStyling}
     key="random1"
     placeholder={"Search for music 'Artist - Track'"}
     value={keyword}
    />
    &nbsp;
    <input
        type='submit'
        style={BtnStyling}
        onSubmit={(e) => setKeyword(e.target.value)}
    />
 </form>
  );
}

export default SearchBar