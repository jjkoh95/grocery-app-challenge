import React from 'react';

const HeaderTab: React.FC = () => {
    return (
        <div style={{ position: 'absolute', zIndex: 9999 }}>
            <header style={{ position: 'relative' }}>
                {/* logo */}
                <div style={{ display: 'table-cell', padding: '20px 0 0 20px' }}>
                </div>
                {/* middle padding */}
                <div style={{ display: 'table-cell', width: '85%' }}>
                    <div style={{ width: '100%', verticalAlign: 'middle' }}></div>
                </div>
                {/* nav tab */}
                <div style={{ display: 'table-cell', paddingRight: '15px' }}>
                    <nav>
                        <ul>
                            <li style={{ display: 'table-cell', padding: '15px' }}>
                                <a href='https://github.com/jjkoh95/grocery-app-challenge'>
                                    <div className='header-tab'>Repo</div>
                                </a>
                            </li>
                            <li style={{ display: 'table-cell', padding: '15px' }}>
                                <a href='mailto:jjkoh95@gmail.com'>
                                    <div className='header-tab'>Contact</div>
                                </a>
                            </li>
                        </ul>
                    </nav>
                </div>
            </header>
        </div>
    );
}

export default HeaderTab;