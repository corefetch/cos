import React, {FC, PropsWithChildren} from "react";
import styled from "styled-components";

const MenuItemSection = styled.div`
    button {

        display: flex;
        align-items: center;
        background: transparent;
        padding-left: 15px;
        border: none;
        width: 100%;
        font-size: 12px;
        text-align: left;
        height: 40px;

        .selected {
            background: #fafafa;
        }

        &:hover {
            background: #fafafa;
            .material-symbols-outlined {
                color: #b6b6b6;
            }
        }

        .material-symbols-outlined {
            margin-right: 5px;
            color: #DDD;
        }
    }
`;

export interface MenuItemProps {
    label: string;
    icon: string;
}

const MenuItem:FC<PropsWithChildren & MenuItemProps> = (props) => {

    return (
        <MenuItemSection>
            <button>
                <span className="material-symbols-outlined">
                    {props.icon}
                </span>
                <label>{props.label}</label>
            </button>
        </MenuItemSection>
    )
}

export interface MenuProps {
    menus: MenuItemProps[]
}

const AsideMenu: FC<PropsWithChildren & MenuProps> = ({menus}) => {
    return (
        <div className="aside-menus">
            {
                menus.map(m => <MenuItem icon={m.icon} label={m.label}></MenuItem>)
            }
            
        </div>
    )
}

export {MenuItem, AsideMenu};