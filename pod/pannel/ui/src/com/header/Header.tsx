import React from 'react';
import styled from 'styled-components';

const HeaderNode = styled.header`
    height: 60px;
    border-bottom: 1px solid var(--main-border-color);
    flex: 1;
`;

const Header:React.FC<React.PropsWithChildren> = ({children}) => {
    return (
        <HeaderNode>
            {children}
        </HeaderNode>
    )
}

export default Header;