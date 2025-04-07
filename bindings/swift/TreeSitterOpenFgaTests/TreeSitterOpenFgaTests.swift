import XCTest
import SwiftTreeSitter
import TreeSitterOpenfga

final class TreeSitterOpenfgaTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_openfga())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading OpenFGA grammar")
    }
}
