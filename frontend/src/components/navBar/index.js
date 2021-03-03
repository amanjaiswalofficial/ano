import React from 'react'
import { Row, Col } from 'react-grid-system';

const NavBar = ({classes}) => {
    return (
        <Row lg={12}>
                    <Col lg={6} sm={6} className={classes.navButton}>
                     Feed goes here
                    </Col>
                    <Col lg={6} sm={6} className={classes.navButton}>
                    Source goes here
                    </Col>
        </Row>
    )
}

export default NavBar
