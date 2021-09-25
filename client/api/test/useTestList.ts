import useSWR from 'swr'
import {Test} from "../../model/test";

export type testListRes = {
    data: Test[];
    isLoading: boolean;
    isError: boolean;
}

export default function useTestList ({initialData}: any) :testListRes {
    const fetcher = () => fetch('http://localhost:18080/test/v1/list').then(res => res.json())
    // @ts-ignore
    const { data, error } = useSWR(`/api/v1/test`, fetcher, {revalidateOnMount: true, initialData})

    return {
        data: data?.data,
        isLoading: !error && !data?.data,
        isError: error
    }
}