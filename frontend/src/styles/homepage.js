import React from 'react'
import {createUseStyles} from 'react-jss'

const useStyles = createUseStyles({
    homepage: {
        width: "700px",
        height: "600px",
    },
    contentRow: {
        minHeight: "555px",
        border: "1px solid red",
        display: "flex"
    },
    navButton: {
        border: "1px solid green"
    }
    // to be used for different screen sizes
    // '@media screen and (max-width: 650px)':{
    //     homepage: {
    //         width: "600px",
    //         height: "600px",
    //         border: "1px solid red"
    //     }
    // }
  })


export default useStyles
