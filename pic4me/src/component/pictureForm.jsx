import React from 'react'

export default props =>
<form  onSubmit={props.onSubmit}>
    <div className="form-group">
        <label htmlFor="titulo">Titulo</label>
        <input type="text"name="title" onChange={props.onChange} 
            className="form-control" id="titulo" value={props.title} 
            placeholder="Title" />
    </div>
    <div className="form-group">
        <label htmlFor="picture">Imagem</label>
        <input type="file" onChange={props.onSelectedFile}  
            className="form-control-file" id="picture"/>
    </div>
    <button type="submit" className="btn btn-primary">Confirmar</button>
</form>