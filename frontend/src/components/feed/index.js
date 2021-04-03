import { Row } from 'react-grid-system';
import NewsItems from "../newsItem"
import React, { useEffect, useState } from 'react';
import { getDataFromAPI } from "../../adapters/feed/adapters"


function Feed({classes}) {
  const [data, setData] = useState([]);

  useEffect(() => {
    async function fetchMyAPI() {
        let response = await getDataFromAPI()
        setData(response)
      }
  
    fetchMyAPI()
  }, [])

  return(
    <Row 
        lg={12} 
        justify="center" 
        className={classes.contentRow}>
            <NewsItems data={data}/>
    </Row>
  )
}

export default Feed;