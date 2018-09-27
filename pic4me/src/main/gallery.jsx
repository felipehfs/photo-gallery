import React from 'react'
import PictureForm from '../component/pictureForm'
import axios from 'axios'
import PictureList from '../component/pictureList'

const URI = "http://localhost:8080/api/photos"
export default class Gallery extends React.Component {
   
    state = {
        filename: '',
        title: "",
        photos: []
    }

    componentDidMount() {
        this.refresh()   
    }

    refresh() {
        axios.get(URI).then(resp => this.setState({...this.state, photos: resp.data}))
    }
    onSelectedFile(e) {
        const {state} = this
        state.filename = e.target.files[0]
        this.setState(state)
    }

    onChange(e) {
        this.setState({...this.state, [e.target.name]:e.target.value })
    }

    onSubmit(e) {
        e.preventDefault()
        const formdata = new FormData()
        formdata.append("title", this.state.title)
        formdata.append("filename", this.state.filename, this.state.filename.name)
        axios.post(URI, formdata, {
            headers:{
                "Content-Type": "multipart/form-data"
            }
        })
        .then(resp => { 
            console.log(resp.data)
            this.setState({...this.state, filename: "", title: ""})
            document.getElementById("picture").value = ""
            this.refresh()
        }).catch(e => console.log(e))
    }

    render() {
        return (
            <div className="container">
                <h1>Photo Gallery</h1>
                <PictureForm {...this.state}
                onChange={this.onChange.bind(this)}
                onSubmit={this.onSubmit.bind(this)}
                onSelectedFile={this.onSelectedFile.bind(this)}></PictureForm>
                <PictureList lista={this.state.photos}/>
            </div>
        )
    }
}