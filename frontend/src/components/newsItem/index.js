import React from 'react'
import {Col,Row, Container} from 'react-grid-system'
import { Divider } from 'antd';

import useStyles from "../../styles/newsItem"

const NewsItem = ({data}) => {

    const classes = useStyles()

    // TODO: Handle exceptions
    const makeLogoURL = ({Source}) => {
        return "https://www.google.com/s2/favicons?sz=256&domain_url=www."+Source.split("/")[2]
    }

    const displayData = () => {

        return data.map((singleItem) => {
            let logoUrl = makeLogoURL(singleItem)
            return (
            <div>
            <Row style={{alignItems: "left"}}>
                <Col lg={1} sm={1} md={1}><img 
                src={logoUrl}
                style={{
                    height:"50px", 
                    marginTop: "10px", 
                    width:"50px" 
                    }}/></Col>
                <Col lg={11} sm={11} md={11} style={{textAlign: "left"}}>
                <a href={singleItem.Link}>
                    <span className={classes.primaryText}>{singleItem.Title} [{singleItem.Date}]</span></a><br/>
                <p className={classes.secondaryText}>{singleItem.Description}</p>
                </Col>
            </Row>
            <Divider style={{height: "20%"}}/>
            </div>)
        })
    }

    return (
        <Container lg={12} style={{width: "100%", overflowY: "scroll", height: "550px"}}>
            {displayData()}
        </Container>
    )
}

export default NewsItem
