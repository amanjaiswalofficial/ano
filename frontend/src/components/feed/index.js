import React from 'react'
import { Row } from 'react-grid-system';

import NewsItem from "../newsItem"

const Feed = ({classes}) => {
    const data = [
        {
            headline: "First Item",
            logoURL: "https://pbs.twimg.com/profile_images/1129666669054324736/1W_E72cn_400x400.png",
            detail: "First Item detail"
        },
        {
            headline: "Second Item",
            logoURL: "https://pbs.twimg.com/profile_images/1129666669054324736/1W_E72cn_400x400.png",
            detail: "Second Item detail"
        },
        {
            headline: "Third Item",
            logoURL: "https://pbs.twimg.com/profile_images/1129666669054324736/1W_E72cn_400x400.png",
            detail: "Third Item detail"
        },
        {
            headline: "Fourth Item",
            logoURL: "https://pbs.twimg.com/profile_images/1129666669054324736/1W_E72cn_400x400.png",
            detail: "Fourth Item detail"
        }
    ]
    return (
        <Row 
        lg={12} 
        justify="center" 
        className={classes.contentRow}>
            <NewsItem data={data}/>
        </Row>
    )
}

export default Feed