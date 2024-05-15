import React, { useState } from 'react';
import styled from 'styled-components';
import Logo from "./logo.svg";

import { AppstoreOutlined, UserSwitchOutlined, MailOutlined, SettingOutlined, DatabaseOutlined, ClusterOutlined, WechatOutlined } from '@ant-design/icons';
import type { MenuProps } from 'antd';
import { Menu } from 'antd';

const AsideWrapper = styled.div`
    width: 200px;
    display: flex;
    flex-direction: column;
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
    overflow: scroll-y;
    flex: 1;
`;

type MenuItem = Required<MenuProps>['items'][number];


const items: MenuItem[] = [
    {
        label: 'Dashboard',
        key: 'gom.dashboard',
        icon: <AppstoreOutlined />,
    },
    {
        label: 'Identity',
        key: 'gom.identity',
        icon: <UserSwitchOutlined />,
        children: [
            { label: 'List', key: 'gom.identity.list' },
            { label: 'Pending', key: 'gom.identity.pending' },
        ]
    },
    {
        label: 'Database',
        key: 'gom.database',
        icon: <DatabaseOutlined />,
        children: [
            { label: 'Schema', key: 'gom.database.schema' },
            { label: 'Data View', key: 'gom.database.view' },
        ]
    },
    {
        label: 'Emails',
        key: 'gom.messages',
        icon: <MailOutlined />,
        children: [
            { label: 'Templates', key: 'gom.messages.templates' },
            { label: 'Campaigns', key: 'gom.messages.campaigns' },
            { label: 'Settings', key: 'gom.messages.settings' },
        ]
    },
    {
        label: 'Support',
        key: 'gom.support',
        icon: <WechatOutlined />,
        children: [
            { label: 'Messages', key: 'gom.support.messages' },
            { label: 'Settings', key: 'gom.support.settings' },
        ]
    },
    {
        label: 'Settings',
        key: 'gom.settings',
        icon: <SettingOutlined/>,
    },
    {
        label: 'System Services',
        key: 'gom.services',
        icon: <ClusterOutlined />,
        children: [
            { label: 'Cloud', key: 'gom.services.cloud' },
            { label: 'Marketplace', key: 'gom.services.marketplace' },
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
            <Separator/>
            <div className="aside-footer">
                Footer
            </div>
        </AsideWrapper>
    )
}

export default Aside;