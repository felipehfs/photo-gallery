import React from 'react'
import "./pictureItem.css"

export default props => 
<div className="card my-card mt-5 mr-5">
  <img className="card-img-top" src={props.src} alt={props.title} />
  <div className="card-body">
    <p className="card-text">{props.title}</p>
  </div>
</div>