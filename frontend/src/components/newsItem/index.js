import React from 'react'
import {Col,Row, Container} from 'react-grid-system'

const NewsItem = ({data}) => {

    const displayData = () => {

        return data.map((singleItem) => {
            return (
            <Row style={{border: "1px solid pink", alignItems: "left"}}>
                <Col lg={1} sm={1} md={1}><img src={singleItem.logoURL} 
                style={{height:"50px", marginTop: "10px", width:"50px" }}></img></Col>
                <Col lg={11} sm={11} md={11} style={{textAlign: "left"}}>
                {singleItem.headline}<br/>
                {singleItem.detail}                
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
