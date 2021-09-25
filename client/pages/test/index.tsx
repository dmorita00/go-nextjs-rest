import Link from 'next/link'
import useTestList from "../../api/test/useTestList";

export default function Test() {

    const { tests, isLoading, isError } = useTestList()
    if (!tests || isLoading || isError) {
        return (
            <>
                loading...
            </>
        )
    }

    return (
    <>
        {tests && tests?.map(test => {
            return (
                <Link href="/" key={test.id}><a>{test.id}</a></Link>
            )
        })}
    </>
  )
}
