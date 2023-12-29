import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';

import 'bootstrap/dist/css/bootstrap.min.css';
import './Layout.css';

const Layout = ({ children }) => {
    const sections = {
        Projects: '/api/projects',
        Containers: '/api/containers',
        Images: '/api/images',
        Networks: '/api/networks'
    };

    return (
        <>
            <Container>
                <Row className="mt-2 mb-4">
                    <Col>
                        <h3>Ahoy</h3>
                    </Col>
                </Row>
                <Row>
                    <Col md={3}>
                        {Object.keys(sections).map((section) => (
                            <div className="fw-bold" key={section}>{section}</div>
                        ))}
                    </Col>
                    <Col md={8}>{children}</Col>
                </Row>
            </Container>
        </>
    );
};

export default Layout;
