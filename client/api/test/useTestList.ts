import useSWR from 'swr'
import {Test} from "../../model/test";

export type testListRes = {
    tests: Test[];
    isLoading: boolean;
    isError: boolean;
}

export default function useTestList () :testListRes {
    const fetcher = () => fetch('http://localhost:18080/test/v1/list').then(res => res.json())
    const { data, error } = useSWR(`/api/v1/test`, fetcher)

    return {
        tests: data?.data,
        isLoading: !error && !data?.data,
        isError: error
    }
}