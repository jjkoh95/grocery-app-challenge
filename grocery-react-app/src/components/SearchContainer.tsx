import React, { ChangeEvent, useEffect } from 'react';

const SearchContainer: React.FC = () => {
    const [search, setSearch] = React.useState<string | undefined>('');
    // const [isLoading, setIsLoading] = React.useState<boolean>(false);
    const [response, setResponse] = React.useState<any[]>([]);

    const [isBrandAsc, setIsBrandAsc] = React.useState<boolean>(false);
    const [isProductAsc, setIsProductAsc] = React.useState<boolean>(false);

    const orderBrand = () => {
        const orderResponse = [...response];
        if (isBrandAsc) {
            setIsBrandAsc(!isBrandAsc);
            orderResponse.sort((a, b) => {
                if (a.brand > b.brand) return -1;
                if (b.brand > a.brand) return 1;
                return 0;
            });
            setResponse(orderResponse);
        } else {
            setIsBrandAsc(!isBrandAsc);
            orderResponse.sort((a, b) => {
                if (a.brand > b.brand) return 1;
                if (b.brand > a.brand) return -1;
                return 0;
            });
            setResponse(orderResponse);
        }
    };

    const orderProduct = () => {
        const orderResponse = [...response];
        if (isProductAsc) {
            setIsProductAsc(!isProductAsc);
            orderResponse.sort((a, b) => {
                if (a.productName > b.productName) return -1;
                if (b.productName > a.productName) return 1;
                return 0;
            });
            setResponse(orderResponse);
        } else {
            setIsProductAsc(!isProductAsc);
            orderResponse.sort((a, b) => {
                if (a.productName > b.productName) return 1;
                if (b.productName > a.productName) return -1;
                return 0;
            });
            setResponse(orderResponse);
        }
    };


    const retrieveResponse = async (searchQuery: string) => {
        const resp = await window.fetch(
            'https://asia-east2-grocery-app-challenge.cloudfunctions.net/GetGrocery',
            {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    searchQuery,
                }),
            }
        ).then(res => res.json());
        setResponse(resp);
    };

    const updateSearch = (e: ChangeEvent<Element>) => {
        const searchQuery: string = (e.target as HTMLInputElement).value
        setSearch(searchQuery);
        retrieveResponse(searchQuery);
    };

    const handleChange = (e: ChangeEvent<Element>, i: number, k: string) => {
        let targetValue: any = (e.target as HTMLInputElement).value;
        if (k === 'UPC12Barcode') {
            // some validation happens here
            targetValue = parseInt(targetValue, 10);
            if (isNaN(targetValue)) return;
        }
        const responseCopy = [...response];
        responseCopy[i][k] = targetValue;
        setResponse(responseCopy);
    }

    const updateRow = async (row: any) => {
        const resp = await window.fetch(
            'https://asia-east2-grocery-app-challenge.cloudfunctions.net/UpsertGrocery',
            {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(row),
            }
        );
        if (resp.status === 200) alert('Successfully updated!');
        else alert('Error!');
    };

    // execute function exactly once
    useEffect(() => {
        retrieveResponse('');
    }, []);

    return (
        <div>
            <div className='input-container'>
                Search by ProductName or Brand: <input type='text' value={search} onChange={updateSearch} />
                <div>
                    <button className='input-button' onClick={orderBrand}>Sort by Brand</button>
                    <button className='input-button' onClick={orderProduct}>Sort by ProductName</button>
                </div>
            </div>
            <div style={{ top: '140px', position: 'relative', marginLeft: '120px' }}>
                {response.map((v, i) => {
                    return (
                        <div key={i}>
                            <input value={v.UPC12Barcode} onChange={(e) => handleChange(e, i, 'UPC12Barcode')} />
                            <input value={v.brand} onChange={(e) => handleChange(e, i, 'brand')} />
                            <input value={v.productName} onChange={(e) => handleChange(e, i, 'productName')} />
                            <button onClick={() => updateRow(v)}>Update</button>
                        </div>
                    );
                })}
            </div>
        </div>
    );
}

export default SearchContainer;