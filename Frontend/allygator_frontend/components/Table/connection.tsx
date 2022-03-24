import React from "react"
import Link from "next/link"

export default function Table(props) {
  console.log(props,"pk")
  const Header = ({ array }) => {
    let counter = 0;
    return array?.map((x) => {
      ++counter;
      return (
        <th className="pl-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider" key={counter}>
          {x.name}
        </th>
      );
    });
  };
  return (
    <>
      <div className="mr-8 flex flex-col">
        <div className="-my-2 overflow-x-auto sm:-mx-6 lg:-mx-8">
          <div className="py-2 align-middle inline-block min-w-full sm:px-6 lg:px-8">
            <div className="shadow overflow-hidden border-b border-gray-200 sm:rounded-lg">
              <table className="table-auto min-w-full divide-y divide-gray-200">
                <thead className="bg-gray-50">
                  <tr>
                    <Header array={props.header} />
                  </tr>
                </thead>
                <tbody className="bg-white divide-y divide-gray-200">
                {props?.data?.map((person) => (
                  <tr key={person.name}>
                    <td className="px-3 py-4 whitespace-nowrap">
                      <div className="flex items-center">
                        <div className="ml-4">
                          <div className="text-sm font-medium text-gray-900">{person?.name}</div>
                          <div className="text-sm text-gray-500">{person.email}</div>
                        </div>
                      </div>
                    </td>

                    <td className="px-3 py-4 whitespace-nowrap text-sm text-gray-500">
                    <button className="flex items-center px-3 py-1 font-semi-bold tracking-wide text-white capitalize transition-colors duration-200 transform bg-blue-500 rounded-md hover:bg-gray-700 focus:outline-none focus:ring focus:ring-blue-300 focus:ring-opacity-80">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    className="h-6 w-6"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M12 9v3m0 0v3m0-3h3m-3 0H9m12 0a9 9 0 11-18 0 9 9 0 0118 0z"
                    />
                  </svg>
                  <span className="mx-1">Connect</span>
                </button>
                    </td>

                  </tr>
                ))}
              </tbody>
              </table>
            </div>
          </div>
        </div>
      </div>
    </>
    

  );
}