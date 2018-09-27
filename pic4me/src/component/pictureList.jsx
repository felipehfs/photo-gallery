import React from 'react'
import PictureItem from './pictureItem'
const URI = "http://localhost:8080/static/"

export default props => {
    if (props.lista.length  > 0){
        return (
                <div className="row">
                    {
                        props.lista.map(elem => <PictureItem key={elem.id} 
                            src={`${URI}/${elem.filename}`} title={elem.title} />)
                }
                </div>
        )
    }  else {
        return <div>Não há dados</div>
    }
}