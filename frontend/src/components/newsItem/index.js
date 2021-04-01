import React from 'react'
import {Col,Row, Container} from 'react-grid-system'

import useStyles from "../../styles/newsItem"

const NewsItem = ({data}) => {

    const classes = useStyles()

    const displayData = () => {

        return data.map((singleItem) => {
            return (
            <Row style={{border: "1px solid pink", alignItems: "left"}}>
                <Col lg={1} sm={1} md={1}><img src="https://www.google.com/s2/favicons?sz=256&domain_url=www.nerdist.com"
                style={{height:"50px", marginTop: "10px", width:"50px" }}></img></Col>
                <Col lg={11} sm={11} md={11} style={{textAlign: "left"}}>
                <a href={singleItem.Link}><span className={classes.primaryText}>{singleItem.Title}</span></a><br/>
                <p className={classes.secondaryText}>{singleItem.Description}</p>
                </Col>
            </Row>)
        })
    }

    return (
        <Container lg={12} style={{border: "1px solid yellow", width: "100%"}}>
            {displayData()}
        </Container>
    )
}

export default NewsItem
