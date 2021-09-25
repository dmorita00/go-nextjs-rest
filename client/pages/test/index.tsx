import Link from 'next/link'
import useTestList from "../../api/test/useTestList";
import {Test} from "../../model/test";

type staticProps = {
    tests: Test[]
}

export async function getStaticProps() {
    const {data} = await fetch('http://server:8080/test/v1/list').then(res => res.json())
    return {
        props: {
            tests: data
        }
    }
}

export default function TestList({tests}: staticProps) {

    const {data, isLoading, isError} = useTestList({initialData: tests})

    if (isLoading) return (<>loading...</>)
    if (!data || isError) return (<>not found!</>)

    return (
    <>
        {data && data?.map(test => {
            return (
                <Link href="/" key={test.id}><a>{test.id}</a></Link>
            )
        })}
    </>
  )
}
