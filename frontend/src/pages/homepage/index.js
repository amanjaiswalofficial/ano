import React from 'react'
import { Container} from 'react-grid-system';
import 'antd/dist/antd.css'

import useStyles from "../../styles/homepage"
import Feed from '../../components/feed';
import NavBar from '../../components/navBar';

const Homepage = () => {

    const classes = useStyles()

    return (
        <div className={classes.homepage}>
        <Container lg={12} fluid="true">
                <Feed 
                 classes={classes}
                />    
                <NavBar 
                 classes={classes}
                />
        </Container>
        </div>
    )
}

export default Homepage
