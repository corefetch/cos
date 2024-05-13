import React, { useState } from 'react';
import styled from 'styled-components';
import Logo from "./logo.svg";

import { AppstoreOutlined, UserSwitchOutlined, MailOutlined, SettingOutlined, DatabaseOutlined, ClusterOutlined, WechatOutlined } from '@ant-design/icons';
import type { MenuProps } from 'antd';
import { Menu } from 'antd';

const AsideWrapper = styled.div`
    width: 200px;
    border-right: 1px solid var(--main-border-color);
`;

const LogoContainer = styled.div`
    height: 50px;
    display: flex;
    align-items: center;
    justify-content: center;
`;

const Separator = styled.div`
    height: 1px;
    width: 100%;
    margin:10px auto;
    background: var(--main-border-color);
`;

const MenuContainer = styled.div`

`;

type MenuItem = Required<MenuProps>['items'][number];


const items: MenuItem[] = [
    {
        label: 'Dashboard',
        key: 'edx.dashboard',
        icon: <AppstoreOutlined />,
    },
    {
        label: 'Accounts',
        key: 'edx.accounts',
        icon: <UserSwitchOutlined />,
        children: [
            { label: 'List', key: 'edx.accounts.list' },
            { label: 'Pending', key: 'edx.accounts.pending' },
        ]
    },
    {
        label: 'Database',
        key: 'edx.database',
        icon: <DatabaseOutlined />,
        children: [
            { label: 'Schema', key: 'edx.database.schema' },
            { label: 'Data View', key: 'edx.database.view' },
        ]
    },
    {
        label: 'Emails',
        key: 'edx.messages',
        icon: <MailOutlined />,
        children: [
            { label: 'Templates', key: 'edx.messages.templates' },
            { label: 'Settings', key: 'edx.messages.settings' },
        ]
    },
    {
        label: 'Support',
        key: 'edx.support',
        icon: <WechatOutlined />,
        children: [
            { label: 'Messages', key: 'edx.support.messages' },
            { label: 'Settings', key: 'edx.support.settings' },
        ]
    },
    {
        label: 'Settings',
        key: 'edx.settings',
        icon: <SettingOutlined/>,
    },
    {
        label: 'System Services',
        key: 'edx.services',
        icon: <ClusterOutlined />,
        children: [
            { label: 'Discovery', key: 'edx.services.discovery' },
            { label: 'Cloud', key: 'edx.services.cloud' },
        ]
    },
];

const Aside: React.FC = () => {

    const [current, setCurrent] = useState('mail');

    const onClick: MenuProps['onClick'] = (e) => {
        console.log('click ', e, current);
        setCurrent(e.key);
    };

    return (
        <AsideWrapper>
            <LogoContainer>
                <a href="/">
                    <img style={{ marginTop: "10px" }} src={Logo} alt="Logo" />
                </a>
            </LogoContainer>
            <Separator />
            <MenuContainer>
                <Menu
                    onClick={onClick}
                    style={{ width: 200, borderInlineEnd: 'none' }}
                    defaultSelectedKeys={['1']}
                    defaultOpenKeys={['sub1']}
                    mode="inline"
                    items={items}
                />
            </MenuContainer>
        </AsideWrapper>
    )
}

export default Aside;