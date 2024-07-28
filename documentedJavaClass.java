package com.demo.util;

// Importing necessary libraries for JSON processing and SQL result
 handling
import org.codehaus.jettison.json.JSONArray;
import org.codehaus.jettison.json.JSONObject;
import java.sql.ResultSet;
// Importing ESAPI for security purposes, though it's not used in
 the provided snippet
import org.owasp.esapi.ESAPI;

/**
* The ToJSON class provides a method to convert a ResultSet from a 
database query
* into a JSONArray. This is useful for creating JSON APIs that 
interact with relational databases.
* It allows for the easy transformation of SQL query results into a
 format that can be easily
* consumed by web applications, mobile applications, or any client 
that understands JSON.
*/
public class ToJSON {

    /**
     * Converts a given ResultSet into a JSONArray.
     * Each row in the ResultSet will be converted into a JSONObject,
     * and each column in the row will be added to the JSONObject with the column name as the key.
     * This method provides a straightforward way to convert relational data into a structured JSON format,
     * making it easier to serialize and transmit data over HTTP or store it in a document-oriented database.
     *
     * @param rs The ResultSet to be converted. This ResultSet is typically obtained from executing
     *           a SQL query against a database.
     * @return JSONArray containing the data from the ResultSet. Each element in the JSONArray
     *         corresponds to a row in the ResultSet, represented as a JSONObject.
     * @throws Exception If there is an error during the conversion process. This could be due to
     *                   issues accessing the ResultSet data, problems with the JSON library, or other
     *                   unexpected issues.
     */
    public JSONArray toJSONArray(ResultSet rs) throws Exception {
        JSONArray json = new JSONArray(); // Initializes the JSON 
array that will be populated and returned
        String temp = null; // Temporary string to hold data, not 
used in the provided snippet

        try {
            // Retrieve metadata from ResultSet which includes 
information about the columns
            java.sql.ResultSetMetaData rsmd = rs.getMetaData();
