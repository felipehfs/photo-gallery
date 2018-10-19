import React from 'react'
import {Switch, Route, Redirect} from 'react-router'
import Gallery from './main/gallery'

export default props =>
<Switch>
    <Route exact path="/" component={Gallery}/>
    <Redirect from="*" to="/" />
</Switch>