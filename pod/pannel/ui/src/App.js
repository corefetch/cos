import './App.css';
import { ConfigProvider } from 'antd';
import Aside from './com/aside/Aside.tsx';
import Header from './com/header/Header.tsx';
import styled from 'styled-components';

const MainWrapper = styled.div`
  display: flex;
  height: 100%;
`;

const BSide = styled.div`
  height: 100%;
  flex: 1;
`;

function App() {
  return (
    <ConfigProvider theme={{ token: { colorPrimary: '#937860' } }}>
      <MainWrapper>
        <Aside />
        <BSide>
          <Header>Header</Header>
        </BSide>
      </MainWrapper>
    </ConfigProvider>
  );
}

export default App;
